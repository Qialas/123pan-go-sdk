package filemgmt

import (
	"context"

	"github.com/Qialas/123pan-go-sdk/core"
)

type MkdirRequest struct {
	Name     string `json:"name"`
	ParentID int64  `json:"parentID"`
}

type MkdirData struct {
	DirID int64 `json:"dirID"`
}

func (s *Service) Mkdir(ctx context.Context, name string, parentID int64) (MkdirData, error) {
	return core.DoAPI[MkdirData](s.client, ctx, "POST", "/upload/v1/file/mkdir", nil, MkdirRequest{Name: name, ParentID: parentID}, true)
}
