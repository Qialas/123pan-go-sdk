package imagebed

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

type ImageFileItem struct {
	FileID         string `json:"fileId"`
	Filename       string `json:"filename"`
	Type           int    `json:"type"`
	Size           int64  `json:"size"`
	Etag           string `json:"etag"`
	Status         int    `json:"status"`
	CreateAt       string `json:"createAt"`
	UpdateAt       string `json:"updateAt"`
	DownloadURL    string `json:"downloadURL"`
	UserSelfURL    string `json:"userSelfURL"`
	TotalTraffic   int64  `json:"totalTraffic"`
	ParentFileID   string `json:"parentFileId"`
	ParentFilename string `json:"parentFilename"`
	Extension      string `json:"extension"`
}

type ListRequest struct {
	ParentFileID string `json:"parentFileId,omitempty"`
	Limit        int    `json:"limit"`
	StartTime    *int64 `json:"startTime,omitempty"`
	EndTime      *int64 `json:"endTime,omitempty"`
	LastFileID   string `json:"lastFileId,omitempty"`
	Type         int    `json:"type"`
}

type ListData struct {
	LastFileID string          `json:"lastFileId"`
	FileList   []ImageFileItem `json:"fileList"`
}

func (s *Service) List(ctx context.Context, req ListRequest) (ListData, error) {
	return core.DoAPI[ListData](s.client, ctx, "POST", "/api/v1/oss/file/list", nil, req, true)
}

func (s *Service) Detail(ctx context.Context, fileID string) (ImageFileItem, error) {
	q := url.Values{}
	q.Set("fileID", fileID)
	return core.DoAPI[ImageFileItem](s.client, ctx, "GET", "/api/v1/oss/file/detail", q, nil, true)
}

type DeleteRequest struct {
	FileIDs []string `json:"fileIDs"`
}

func (s *Service) Delete(ctx context.Context, fileIDs []string) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "POST", "/api/v1/oss/file/delete", nil, DeleteRequest{FileIDs: fileIDs}, true)
	return err
}

type MoveRequest struct {
	FileIDs        []string `json:"fileIDs"`
	ToParentFileID string   `json:"toParentFileID"`
}

func (s *Service) Move(ctx context.Context, fileIDs []string, toParentFileID string) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "POST", "/api/v1/oss/file/move", nil, MoveRequest{FileIDs: fileIDs, ToParentFileID: toParentFileID}, true)
	return err
}

type SourceCopyRequest struct {
	FileIDs        []string `json:"fileIDs"`
	ToParentFileID string   `json:"toParentFileID,omitempty"`
	SourceType     int      `json:"sourceType"`
	Type           int      `json:"type"`
}

type SourceCopyData struct {
	TaskID string `json:"taskID"`
}

func (s *Service) SourceCopy(ctx context.Context, req SourceCopyRequest) (SourceCopyData, error) {
	return core.DoAPI[SourceCopyData](s.client, ctx, "POST", "/api/v1/oss/source/copy", nil, req, true)
}

type SourceCopyProcessData struct {
	Status  int    `json:"status"`
	FailMsg string `json:"failMsg"`
}

func (s *Service) SourceCopyProcess(ctx context.Context, taskID string) (SourceCopyProcessData, error) {
	q := url.Values{}
	q.Set("taskID", taskID)
	return core.DoAPI[SourceCopyProcessData](s.client, ctx, "GET", "/api/v1/oss/source/copy/process", q, nil, true)
}

type SourceCopyFailItem struct {
	FileID   int64  `json:"fileId"`
	Filename string `json:"filename"`
}

type SourceCopyFailData struct {
	Total int64                `json:"total"`
	List  []SourceCopyFailItem `json:"list"`
}

type SourceCopyFailQuery struct {
	TaskID string
	Limit  int
	Page   int
}

func (s *Service) SourceCopyFail(ctx context.Context, q SourceCopyFailQuery) (SourceCopyFailData, error) {
	qs := url.Values{}
	qs.Set("taskID", q.TaskID)
	qs.Set("limit", strconv.Itoa(q.Limit))
	qs.Set("page", strconv.Itoa(q.Page))
	return core.DoAPI[SourceCopyFailData](s.client, ctx, "GET", "/api/v1/oss/source/copy/fail", qs, nil, true)
}

type OfflineMigrateRequest struct {
	URL           string `json:"url"`
	FileName      string `json:"fileName,omitempty"`
	BusinessDirID string `json:"businessDirID,omitempty"`
	CallBackURL   string `json:"callBackUrl,omitempty"`
	Type          int    `json:"type"`
}

type OfflineMigrateData struct {
	TaskID int64 `json:"taskID"`
}

func (s *Service) OfflineMigrate(ctx context.Context, req OfflineMigrateRequest) (OfflineMigrateData, error) {
	return core.DoAPI[OfflineMigrateData](s.client, ctx, "POST", "/api/v1/oss/offline/download", nil, req, true)
}

type OfflineMigrateProcessData struct {
	Status  int `json:"status"`
	Process int `json:"process"`
}

func (s *Service) OfflineMigrateProcess(ctx context.Context, taskID int64) (OfflineMigrateProcessData, error) {
	q := url.Values{}
	q.Set("taskID", strconv.FormatInt(taskID, 10))
	return core.DoAPI[OfflineMigrateProcessData](s.client, ctx, "GET", "/api/v1/oss/offline/download/process", q, nil, true)
}
