package filemgmt

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Qialas/123pan-go-sdk/core"
)

type DownloadInfoData struct {
	DownloadURL string `json:"downloadUrl"`
}

func (s *Service) DownloadInfo(ctx context.Context, fileID int64) (DownloadInfoData, error) {
	q := url.Values{}
	q.Set("fileId", strconv.FormatInt(fileID, 10))
	return core.DoAPI[DownloadInfoData](s.client, ctx, "GET", "/api/v1/file/download_info", q, nil, true)
}
