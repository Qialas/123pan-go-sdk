package core

import (
	"errors"
	"net/http"
	"net/url"
)

type Option func(*Client) error

func WithAccessToken(token string) Option {
	return func(c *Client) error {
		c.accessToken = token
		return nil
	}
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) error {
		if hc == nil {
			return errors.New("http client is nil")
		}
		c.httpClient = hc
		return nil
	}
}

func WithBaseURL(raw string) Option {
	return func(c *Client) error {
		u, err := url.Parse(raw)
		if err != nil {
			return err
		}
		c.baseURL = u
		return nil
	}
}

func WithPlatform(platform string) Option {
	return func(c *Client) error {
		c.platform = platform
		return nil
	}
}
