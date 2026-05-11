package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var errBaseURLNil = errors.New("base url is nil")

func (c *Client) resolveURL(path string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	if u.IsAbs() {
		return u.String(), nil
	}
	if c.baseURL == nil {
		return "", &url.Error{Op: "resolve", URL: path, Err: errBaseURLNil}
	}
	return c.baseURL.ResolveReference(u).String(), nil
}

func (c *Client) newRequest(ctx context.Context, method, fullURL string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Platform", c.platform)
	return req, nil
}

func (c *Client) do(req *http.Request) (*http.Response, []byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}
	return resp, b, nil
}

func (c *Client) doJSON(ctx context.Context, method, path string, query url.Values, reqBody any, out any, withAuth bool) error {
	fullURL, err := c.resolveURL(path)
	if err != nil {
		return err
	}
	if query != nil {
		u, err := url.Parse(fullURL)
		if err != nil {
			return err
		}
		u.RawQuery = query.Encode()
		fullURL = u.String()
	}

	var body io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	req, err := c.newRequest(ctx, method, fullURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if withAuth {
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
	}

	resp, b, err := c.do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var r APIResponse[json.RawMessage]
		_ = json.Unmarshal(b, &r)
		return &APIError{StatusCode: resp.StatusCode, Code: r.Code, Message: r.Message, TraceID: r.TraceID}
	}

	if out == nil {
		var r APIResponse[json.RawMessage]
		if err := json.Unmarshal(b, &r); err != nil {
			return err
		}
		if r.Code != 0 {
			return &APIError{StatusCode: resp.StatusCode, Code: r.Code, Message: r.Message, TraceID: r.TraceID}
		}
		return nil
	}

	if err := json.Unmarshal(b, out); err != nil {
		return err
	}
	return nil
}

func DoAPI[T any](c *Client, ctx context.Context, method, path string, query url.Values, reqBody any, withAuth bool) (T, error) {
	var zero T
	var resp APIResponse[T]
	if err := c.doJSON(ctx, method, path, query, reqBody, &resp, withAuth); err != nil {
		return zero, err
	}
	if resp.Code != 0 {
		return zero, &APIError{StatusCode: 200, Code: resp.Code, Message: resp.Message, TraceID: resp.TraceID}
	}
	return resp.Data, nil
}
