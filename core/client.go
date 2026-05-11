package core

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	accessToken string
	platform    string
}

func NewClient(opts ...Option) (*Client, error) {
	u, _ := url.Parse("https://open-api.123pan.com")
	c := &Client{
		baseURL:  u,
		platform: "open_platform",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) BaseURL() string {
	if c.baseURL == nil {
		return ""
	}
	return c.baseURL.String()
}

func (c *Client) AccessToken() string {
	return c.accessToken
}

func (c *Client) SetAccessToken(token string) {
	c.accessToken = token
}

func (c *Client) Platform() string {
	return c.platform
}

func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}
