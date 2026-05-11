package filemgmt

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/Qialas/123pan-go-sdk/core"
)

type CreateFileV2Request struct {
	ParentFileID int64  `json:"parentFileID"`
	Filename     string `json:"filename"`
	Etag         string `json:"etag"`
	Size         int64  `json:"size"`
	Duplicate    *int   `json:"duplicate,omitempty"`
	ContainDir   *bool  `json:"containDir,omitempty"`
}

type CreateFileV2Data struct {
	FileID      int64    `json:"fileID"`
	PreuploadID string   `json:"preuploadID,omitempty"`
	Reuse       bool     `json:"reuse"`
	SliceSize   int64    `json:"sliceSize"`
	Servers     []string `json:"servers"`
}

func (s *Service) CreateFileV2(ctx context.Context, req CreateFileV2Request) (CreateFileV2Data, error) {
	return core.DoAPI[CreateFileV2Data](s.client, ctx, "POST", "/upload/v2/file/create", nil, req, true)
}

func (s *Service) UploadDomains(ctx context.Context) ([]string, error) {
	return core.DoAPI[[]string](s.client, ctx, "GET", "/upload/v2/file/domain", nil, nil, true)
}

type UploadCompleteRequest struct {
	PreuploadID string `json:"preuploadID"`
}

type UploadCompleteData struct {
	Completed bool  `json:"completed"`
	FileID    int64 `json:"fileID"`
}

func (s *Service) UploadComplete(ctx context.Context, preuploadID string) (UploadCompleteData, error) {
	return core.DoAPI[UploadCompleteData](s.client, ctx, "POST", "/upload/v2/file/upload_complete", nil, UploadCompleteRequest{PreuploadID: preuploadID}, true)
}

func (s *Service) UploadSlice(ctx context.Context, uploadBase, preuploadID string, sliceNo int, sliceMD5 string, filename string, sliceReader io.Reader) error {
	base, err := url.Parse(uploadBase)
	if err != nil {
		return err
	}
	base.Path = path.Join(base.Path, "/upload/v2/file/slice")
	fullURL := base.String()

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer mw.Close()

		_ = mw.WriteField("preuploadID", preuploadID)
		_ = mw.WriteField("sliceNo", strconv.Itoa(sliceNo))
		_ = mw.WriteField("sliceMD5", sliceMD5)

		part, err := mw.CreateFormFile("slice", filename)
		if err != nil {
			_ = pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, sliceReader); err != nil {
			_ = pw.CloseWithError(err)
			return
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, pr)
	if err != nil {
		return err
	}
	req.Header.Set("Platform", s.client.Platform())
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+s.client.AccessToken())

	hc := s.client.HTTPClient()
	if hc == nil {
		hc = http.DefaultClient
	}
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var apiResp core.APIResponse[json.RawMessage]
		_ = json.NewDecoder(resp.Body).Decode(&apiResp)
		return &core.APIError{StatusCode: resp.StatusCode, Code: apiResp.Code, Message: apiResp.Message, TraceID: apiResp.TraceID}
	}
	var apiResp core.APIResponse[struct{}]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return err
	}
	if apiResp.Code != 0 {
		return &core.APIError{StatusCode: resp.StatusCode, Code: apiResp.Code, Message: apiResp.Message, TraceID: apiResp.TraceID}
	}
	return nil
}

func (s *Service) UploadSliceBytes(ctx context.Context, uploadBase, preuploadID string, sliceNo int, filename string, b []byte) error {
	return s.UploadSlice(ctx, uploadBase, preuploadID, sliceNo, MD5Hex(b), filename, bytes.NewReader(b))
}

func MD5Hex(b []byte) string {
	h := md5.Sum(b)
	return hex.EncodeToString(h[:])
}

func pickUploadBase(servers []string) string {
	for _, s := range servers {
		s = strings.TrimSpace(s)
		if s != "" {
			return strings.TrimRight(s, "/")
		}
	}
	return ""
}

type UploadFileOptions struct {
	Duplicate  *int
	ContainDir *bool
}

func (s *Service) UploadFileFromPath(ctx context.Context, parentFileID int64, localPath string, opt *UploadFileOptions) (int64, error) {
	f, err := os.Open(localPath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil {
		return 0, err
	}

	sum, err := md5File(f)
	if err != nil {
		return 0, err
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return 0, err
	}

	req := CreateFileV2Request{
		ParentFileID: parentFileID,
		Filename:     st.Name(),
		Etag:         sum,
		Size:         st.Size(),
	}
	if opt != nil {
		req.Duplicate = opt.Duplicate
		req.ContainDir = opt.ContainDir
	}

	create, err := s.CreateFileV2(ctx, req)
	if err != nil {
		return 0, err
	}
	if create.Reuse {
		if create.FileID == 0 {
			return 0, errors.New("reuse=true but fileID is empty")
		}
		return create.FileID, nil
	}

	uploadBase := pickUploadBase(create.Servers)
	if uploadBase == "" {
		domains, err := s.UploadDomains(ctx)
		if err != nil {
			return 0, err
		}
		uploadBase = pickUploadBase(domains)
	}
	if uploadBase == "" {
		return 0, errors.New("upload server is empty")
	}

	buf := make([]byte, create.SliceSize)
	sliceNo := 1
	for {
		n, err := io.ReadFull(f, buf)
		if errors.Is(err, io.ErrUnexpectedEOF) {
			if n == 0 {
				break
			}
			if err := s.UploadSliceBytes(ctx, uploadBase, create.PreuploadID, sliceNo, st.Name(), buf[:n]); err != nil {
				return 0, err
			}
			sliceNo++
			break
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return 0, err
		}
		if err := s.UploadSliceBytes(ctx, uploadBase, create.PreuploadID, sliceNo, st.Name(), buf[:n]); err != nil {
			return 0, err
		}
		sliceNo++
	}

	complete, err := s.UploadComplete(ctx, create.PreuploadID)
	if err != nil {
		return 0, err
	}
	if !complete.Completed {
		return 0, errors.New("upload not completed")
	}
	return complete.FileID, nil
}
