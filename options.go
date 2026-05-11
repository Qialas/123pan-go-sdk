package pan123

import (
	"net/http"

	"github.com/Qialas/123pan-go-sdk/core"
)

type Option = core.Option

func WithAccessToken(token string) Option {
	return core.WithAccessToken(token)
}

func WithHTTPClient(hc *http.Client) Option {
	return core.WithHTTPClient(hc)
}

func WithBaseURL(raw string) Option {
	return core.WithBaseURL(raw)
}

func WithPlatform(platform string) Option {
	return core.WithPlatform(platform)
}
