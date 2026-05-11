package filemgmt

import (
	"context"

	"github.com/Qialas/123pan-go-sdk/core"
)

type CopyFileRequest struct {
	FileID      int64 `json:"fileId"`
	TargetDirID int64 `json:"targetDirId"`
}

type CopyFileData struct {
	SourceFileID int64 `json:"sourceFileId"`
	TargetFileID int64 `json:"targetFileId"`
}

func (s *Service) Copy(ctx context.Context, fileID, targetDirID int64) (CopyFileData, error) {
	return core.DoAPI[CopyFileData](s.client, ctx, "POST", "/api/v1/file/copy", nil, CopyFileRequest{FileID: fileID, TargetDirID: targetDirID}, true)
}

type AsyncCopyRequest struct {
	FileIDs     []int64 `json:"fileIds"`
	TargetDirID int64   `json:"targetDirId"`
}

type AsyncCopyData struct {
	TaskID int64 `json:"taskId"`
}

func (s *Service) AsyncCopy(ctx context.Context, fileIDs []int64, targetDirID int64) (AsyncCopyData, error) {
	return core.DoAPI[AsyncCopyData](s.client, ctx, "POST", "/api/v1/file/async/copy", nil, AsyncCopyRequest{FileIDs: fileIDs, TargetDirID: targetDirID}, true)
}

type AsyncCopyProcessRequest struct {
	TaskID int64 `json:"taskId"`
}

type AsyncCopyProcessData struct {
	TaskID int64 `json:"taskId"`
	Status int   `json:"status"`
}

func (s *Service) AsyncCopyProcess(ctx context.Context, taskID int64) (AsyncCopyProcessData, error) {
	return core.DoAPI[AsyncCopyProcessData](s.client, ctx, "GET", "/api/v1/file/async/copy/process", nil, AsyncCopyProcessRequest{TaskID: taskID}, true)
}
