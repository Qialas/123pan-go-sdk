package pan123

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Qialas/123pan-go-sdk/usermgmt"
)

func TestGetAccessToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/access_token" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Header.Get("Platform") != "open_platform" {
			t.Fatalf("platform header: %q", r.Header.Get("Platform"))
		}
		var req usermgmt.AccessTokenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}
		_ = json.NewEncoder(w).Encode(APIResponse[AccessTokenData]{
			Code:    0,
			Message: "ok",
			Data:    AccessTokenData{AccessToken: "t", ExpiredAt: "2099-01-01 00:00:00"},
			TraceID: "",
		})
	}))
	defer ts.Close()

	c, err := NewClient(WithBaseURL(ts.URL))
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.GetAccessToken(context.Background(), "a", "b")
	if err != nil {
		t.Fatal(err)
	}
	if got.AccessToken != "t" {
		t.Fatalf("access token: %q", got.AccessToken)
	}
}

func TestUserInfoAuth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/user/info" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Header.Get("Authorization") != "Bearer t" {
			_ = json.NewEncoder(w).Encode(APIResponse[any]{Code: 401, Message: "no", Data: nil, TraceID: "x"})
			return
		}
		_ = json.NewEncoder(w).Encode(APIResponse[UserInfoData]{Code: 0, Message: "ok", Data: UserInfoData{UID: 1, Nickname: "n"}, TraceID: ""})
	}))
	defer ts.Close()

	c, err := NewClient(WithBaseURL(ts.URL), WithAccessToken("t"))
	if err != nil {
		t.Fatal(err)
	}
	u, err := c.User.Info(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if u.UID != 1 {
		t.Fatalf("uid: %d", u.UID)
	}
}

func TestAPIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(APIResponse[any]{Code: 401, Message: "access_token无效", Data: nil, TraceID: "trace"})
	}))
	defer ts.Close()

	c, err := NewClient(WithBaseURL(ts.URL), WithAccessToken("bad"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.User.Info(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	if _, ok := err.(*APIError); !ok {
		t.Fatalf("expected APIError, got %T", err)
	}
}
