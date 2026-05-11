package filemgmt

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Qialas/123pan-go-sdk/core"
)

type FileDetailData struct {
	FileID       int64  `json:"fileID"`
	Filename     string `json:"filename"`
	Type         int    `json:"type"`
	Size         int64  `json:"size"`
	Etag         string `json:"etag"`
	Status       int    `json:"status"`
	ParentFileID int64  `json:"parentFileID"`
	CreateAt     string `json:"createAt"`
	Trashed      int    `json:"trashed"`
}

func (s *Service) Detail(ctx context.Context, fileID int64) (FileDetailData, error) {
	q := url.Values{}
	q.Set("fileID", strconv.FormatInt(fileID, 10))
	return core.DoAPI[FileDetailData](s.client, ctx, "GET", "/api/v1/file/detail", q, nil, true)
}

type InfosRequest struct {
	FileIDs []int64 `json:"fileIds"`
}

type InfosItem struct {
	FileID       int64  `json:"fileId"`
	Filename     string `json:"filename"`
	ParentFileID int64  `json:"parentFileId"`
	Type         int    `json:"type"`
	Etag         string `json:"etag"`
	Size         int64  `json:"size"`
	Category     int    `json:"category"`
	Status       int    `json:"status"`
	PunishFlag   int    `json:"punishFlag"`
	S3KeyFlag    string `json:"s3KeyFlag"`
	StorageNode  string `json:"storageNode"`
	Trashed      int    `json:"trashed"`
	CreateAt     string `json:"createAt"`
	UpdateAt     int64  `json:"updateAt"`
}

type InfosData struct {
	FileList []InfosItem `json:"fileList"`
}

func (s *Service) Infos(ctx context.Context, fileIDs []int64) (InfosData, error) {
	return core.DoAPI[InfosData](s.client, ctx, "POST", "/api/v1/file/infos", nil, InfosRequest{FileIDs: fileIDs}, true)
}
