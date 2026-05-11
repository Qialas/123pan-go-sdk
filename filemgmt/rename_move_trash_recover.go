package filemgmt

import (
	"context"

	"github.com/Qialas/123pan-go-sdk/core"
)

type RenameRequest struct {
	FileID   int64  `json:"fileId"`
	FileName string `json:"fileName"`
}

func (s *Service) Rename(ctx context.Context, fileID int64, fileName string) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "PUT", "/api/v1/file/name", nil, RenameRequest{FileID: fileID, FileName: fileName}, true)
	return err
}

type BatchRenameRequest struct {
	RenameList []string `json:"renameList"`
}

type BatchRenameSuccessItem struct {
	FileID   int64  `json:"fileID"`
	UpdateAt string `json:"updateAt"`
}

type BatchRenameFailItem struct {
	FileID  int64  `json:"fileID"`
	Message string `json:"message"`
}

type BatchRenameData struct {
	SuccessList []BatchRenameSuccessItem `json:"successList"`
	FailList    []BatchRenameFailItem    `json:"failList"`
}

func (s *Service) BatchRename(ctx context.Context, renameList []string) (BatchRenameData, error) {
	return core.DoAPI[BatchRenameData](s.client, ctx, "POST", "/api/v1/file/rename", nil, BatchRenameRequest{RenameList: renameList}, true)
}

type MoveRequest struct {
	FileIDs        []int64 `json:"fileIDs"`
	ToParentFileID int64   `json:"toParentFileID"`
}

func (s *Service) Move(ctx context.Context, fileIDs []int64, toParentFileID int64) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "POST", "/api/v1/file/move", nil, MoveRequest{FileIDs: fileIDs, ToParentFileID: toParentFileID}, true)
	return err
}

type TrashRequest struct {
	FileIDs []int64 `json:"fileIDs"`
}

func (s *Service) Trash(ctx context.Context, fileIDs []int64) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "POST", "/api/v1/file/trash", nil, TrashRequest{FileIDs: fileIDs}, true)
	return err
}

type RecoverRequest struct {
	FileIDs []int64 `json:"fileIDs"`
}

type RecoverData struct {
	AbnormalFileIDs []int64 `json:"abnormalFileIDs"`
}

func (s *Service) Recover(ctx context.Context, fileIDs []int64) (RecoverData, error) {
	return core.DoAPI[RecoverData](s.client, ctx, "POST", "/api/v1/file/recover", nil, RecoverRequest{FileIDs: fileIDs}, true)
}

type RecoverByPathRequest struct {
	FileIDs      []int64 `json:"fileIDs"`
	ParentFileID int64   `json:"parentFileID"`
}

func (s *Service) RecoverByPath(ctx context.Context, fileIDs []int64, parentFileID int64) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "POST", "/api/v1/file/recover/by_path", nil, RecoverByPathRequest{FileIDs: fileIDs, ParentFileID: parentFileID}, true)
	return err
}
