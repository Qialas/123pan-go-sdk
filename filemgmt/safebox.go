package filemgmt

import (
	"context"
	"net/url"

	"github.com/Qialas/123pan-go-sdk/core"
)

type SafeBoxIDData struct {
	FileID int64 `json:"fileId"`
}

func (s *Service) SafeBoxID(ctx context.Context, password string) (SafeBoxIDData, error) {
	q := url.Values{}
	q.Set("password", password)
	return core.DoAPI[SafeBoxIDData](s.client, ctx, "GET", "/api/v1/safebox/id", q, nil, true)
}
