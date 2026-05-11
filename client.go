package pan123

import (
	"context"

	"github.com/Qialas/123pan-go-sdk/core"
	"github.com/Qialas/123pan-go-sdk/directlinkmgmt"
	"github.com/Qialas/123pan-go-sdk/filemgmt"
	"github.com/Qialas/123pan-go-sdk/imagebed"
	"github.com/Qialas/123pan-go-sdk/offlinedl"
	"github.com/Qialas/123pan-go-sdk/sharemgmt"
	"github.com/Qialas/123pan-go-sdk/usermgmt"
)

type Client struct {
	*core.Client

	File       *filemgmt.Service
	Share      *sharemgmt.Service
	Offline    *offlinedl.Service
	User       *usermgmt.Service
	DirectLink *directlinkmgmt.Service
	Image      *imagebed.Service
}

func NewClient(opts ...Option) (*Client, error) {
	base, err := core.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	c := &Client{Client: base}
	c.File = filemgmt.New(base)
	c.Share = sharemgmt.New(base)
	c.Offline = offlinedl.New(base)
	c.User = usermgmt.New(base)
	c.DirectLink = directlinkmgmt.New(base)
	c.Image = imagebed.New(base)
	return c, nil
}

func (c *Client) Ping(ctx context.Context) error {
	_, err := c.User.Info(ctx)
	return err
}

func (c *Client) GetAccessToken(ctx context.Context, clientID, clientSecret string) (usermgmt.AccessTokenData, error) {
	return c.User.GetAccessToken(ctx, clientID, clientSecret)
}
