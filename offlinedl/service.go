package offlinedl

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Qialas/123pan-go-sdk/core"
)

type Service struct {
	client *core.Client
}

func New(c *core.Client) *Service {
	return &Service{client: c}
}

type CreateDownloadRequest struct {
	URL         string `json:"url"`
	FileName    string `json:"fileName,omitempty"`
	DirID       *int64 `json:"dirID,omitempty"`
	CallBackURL string `json:"callBackUrl,omitempty"`
}

type CreateDownloadData struct {
	TaskID int64 `json:"taskID"`
}

func (s *Service) CreateDownload(ctx context.Context, req CreateDownloadRequest) (CreateDownloadData, error) {
	return core.DoAPI[CreateDownloadData](s.client, ctx, "POST", "/api/v1/offline/download", nil, req, true)
}

type DownloadProcessData struct {
	Process float64 `json:"process"`
	Status  int     `json:"status"`
}

func (s *Service) DownloadProcess(ctx context.Context, taskID int64) (DownloadProcessData, error) {
	q := url.Values{}
	q.Set("taskID", strconv.FormatInt(taskID, 10))
	return core.DoAPI[DownloadProcessData](s.client, ctx, "GET", "/api/v1/offline/download/process", q, nil, true)
}
