package filemgmt

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Qialas/123pan-go-sdk/core"
)

type FileListV2Item struct {
	FileID       int64  `json:"fileId"`
	Filename     string `json:"filename"`
	Type         int    `json:"type"`
	Size         int64  `json:"size"`
	Etag         string `json:"etag"`
	Status       int    `json:"status"`
	ParentFileID int64  `json:"parentFileId"`
	Category     int    `json:"category"`
	Trashed      int    `json:"trashed"`
}

type FileListV2Data struct {
	LastFileID int64            `json:"lastFileId"`
	FileList   []FileListV2Item `json:"fileList"`
}

type ListV2Request struct {
	ParentFileID int64
	Limit        int
	SearchData   string
	SearchMode   *int
	LastFileID   *int64
}

func (s *Service) ListV2(ctx context.Context, req ListV2Request) (FileListV2Data, error) {
	q := url.Values{}
	q.Set("parentFileId", strconv.FormatInt(req.ParentFileID, 10))
	q.Set("limit", strconv.Itoa(req.Limit))
	if req.SearchData != "" {
		q.Set("searchData", req.SearchData)
	}
	if req.SearchMode != nil {
		q.Set("searchMode", strconv.Itoa(*req.SearchMode))
	}
	if req.LastFileID != nil {
		q.Set("lastFileId", strconv.FormatInt(*req.LastFileID, 10))
	}
	return core.DoAPI[FileListV2Data](s.client, ctx, "GET", "/api/v2/file/list", q, nil, true)
}

type FileListV1Item struct {
	FileID       int64  `json:"fileID"`
	Filename     string `json:"filename"`
	Type         int    `json:"type"`
	Size         int64  `json:"size"`
	Etag         bool   `json:"etag"`
	Status       int    `json:"status"`
	ParentFileID int64  `json:"parentFileId"`
	ParentName   string `json:"parentName"`
	Category     int    `json:"category"`
	ContentType  string `json:"contentType"`
}

type FileListV1Data struct {
	FileList []FileListV1Item `json:"fileList"`
}

type ListV1Request struct {
	ParentFileID   int64
	Page           int
	Limit          int
	OrderBy        string
	OrderDirection string
	Trashed        *bool
	SearchData     string
}

func (s *Service) ListV1(ctx context.Context, req ListV1Request) (FileListV1Data, error) {
	q := url.Values{}
	q.Set("parentFileId", strconv.FormatInt(req.ParentFileID, 10))
	q.Set("page", strconv.Itoa(req.Page))
	q.Set("limit", strconv.Itoa(req.Limit))
	q.Set("orderBy", req.OrderBy)
	q.Set("orderDirection", req.OrderDirection)
	if req.Trashed != nil {
		q.Set("trashed", strconv.FormatBool(*req.Trashed))
	}
	if req.SearchData != "" {
		q.Set("searchData", req.SearchData)
	}
	return core.DoAPI[FileListV1Data](s.client, ctx, "GET", "/api/v1/file/list", q, nil, true)
}
