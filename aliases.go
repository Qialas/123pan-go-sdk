package pan123

import (
	"github.com/Qialas/123pan-go-sdk/core"
	"github.com/Qialas/123pan-go-sdk/directlinkmgmt"
	"github.com/Qialas/123pan-go-sdk/filemgmt"
	"github.com/Qialas/123pan-go-sdk/imagebed"
	"github.com/Qialas/123pan-go-sdk/offlinedl"
	"github.com/Qialas/123pan-go-sdk/sharemgmt"
	"github.com/Qialas/123pan-go-sdk/usermgmt"
)

type APIError = core.APIError

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	TraceID string `json:"x-traceID"`
}

type AccessTokenData = usermgmt.AccessTokenData

type UserInfoData = usermgmt.UserInfoData

type FileListV2Data = filemgmt.FileListV2Data
type FileListV2Item = filemgmt.FileListV2Item
type FileDetailData = filemgmt.FileDetailData

type CreateFileV2Request = filemgmt.CreateFileV2Request
type CreateFileV2Data = filemgmt.CreateFileV2Data
type UploadCompleteData = filemgmt.UploadCompleteData
type UploadFileOptions = filemgmt.UploadFileOptions

type DownloadInfoData = filemgmt.DownloadInfoData

type CreateShareRequest = sharemgmt.CreateShareRequest
type CreateShareData = sharemgmt.CreateShareData
type ShareListData = sharemgmt.ShareListData
type UpdateShareRequest = sharemgmt.UpdateShareRequest
type CreatePaidShareRequest = sharemgmt.CreatePaidShareRequest
type PaidShareListData = sharemgmt.PaidShareListData
type UpdatePaidShareRequest = sharemgmt.UpdatePaidShareRequest

type CreateOfflineDownloadRequest = offlinedl.CreateDownloadRequest
type CreateOfflineDownloadData = offlinedl.CreateDownloadData
type OfflineDownloadProcessData = offlinedl.DownloadProcessData

type DirectLinkURLData = directlinkmgmt.URLData

type SafeBoxIDData = filemgmt.SafeBoxIDData

type ImageFileItem = imagebed.ImageFileItem
