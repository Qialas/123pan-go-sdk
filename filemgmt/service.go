package filemgmt

import "github.com/Qialas/123pan-go-sdk/core"

type Service struct {
	client *core.Client
}

func New(c *core.Client) *Service {
	return &Service{client: c}
}
