package filemgmt

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/Qialas/123pan-go-sdk/core"
)

type Sha1ReuseRequest struct {
	ParentFileID int64  `json:"parentFileID"`
	Filename     string `json:"filename"`
	Sha1         string `json:"sha1"`
	Size         int64  `json:"size"`
	Duplicate    *int   `json:"duplicate,omitempty"`
}

type Sha1ReuseData struct {
	FileID int64 `json:"fileID"`
	Reuse  bool  `json:"reuse"`
}

func (s *Service) Sha1Reuse(ctx context.Context, req Sha1ReuseRequest) (Sha1ReuseData, error) {
	return core.DoAPI[Sha1ReuseData](s.client, ctx, "POST", "/upload/v2/file/sha1_reuse", nil, req, true)
}

type SingleUploadRequest struct {
	ParentFileID int64
	Filename     string
	Etag         string
	Size         int64
	Duplicate    *int
	ContainDir   *bool
}

type SingleUploadData struct {
	FileID    int64 `json:"fileID"`
	Completed bool  `json:"completed"`
}

func (s *Service) SingleUpload(ctx context.Context, uploadBase string, req SingleUploadRequest, fileFieldName string, r io.Reader) (SingleUploadData, error) {
	base, err := url.Parse(uploadBase)
	if err != nil {
		return SingleUploadData{}, err
	}
	base.Path = path.Join(base.Path, "/upload/v2/file/single/create")
	fullURL := base.String()

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer mw.Close()

		_ = mw.WriteField("parentFileID", strconv.FormatInt(req.ParentFileID, 10))
		_ = mw.WriteField("filename", req.Filename)
		_ = mw.WriteField("etag", req.Etag)
		_ = mw.WriteField("size", strconv.FormatInt(req.Size, 10))
		if req.Duplicate != nil {
			_ = mw.WriteField("duplicate", strconv.Itoa(*req.Duplicate))
		}
		if req.ContainDir != nil {
			_ = mw.WriteField("containDir", strconv.FormatBool(*req.ContainDir))
		}
		part, err := mw.CreateFormFile("file", fileFieldName)
		if err != nil {
			_ = pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, r); err != nil {
			_ = pw.CloseWithError(err)
			return
		}
	}()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, pr)
	if err != nil {
		return SingleUploadData{}, err
	}
	httpReq.Header.Set("Platform", s.client.Platform())
	httpReq.Header.Set("Content-Type", mw.FormDataContentType())
	httpReq.Header.Set("Authorization", "Bearer "+s.client.AccessToken())

	hc := s.client.HTTPClient()
	if hc == nil {
		hc = http.DefaultClient
	}
	resp, err := hc.Do(httpReq)
	if err != nil {
		return SingleUploadData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var apiResp core.APIResponse[json.RawMessage]
		_ = json.NewDecoder(resp.Body).Decode(&apiResp)
		return SingleUploadData{}, &core.APIError{StatusCode: resp.StatusCode, Code: apiResp.Code, Message: apiResp.Message, TraceID: apiResp.TraceID}
	}
	var apiResp core.APIResponse[SingleUploadData]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return SingleUploadData{}, err
	}
	if apiResp.Code != 0 {
		return SingleUploadData{}, &core.APIError{StatusCode: resp.StatusCode, Code: apiResp.Code, Message: apiResp.Message, TraceID: apiResp.TraceID}
	}
	return apiResp.Data, nil
}
