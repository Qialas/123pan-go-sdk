package directlinkmgmt

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

type URLData struct {
	URL string `json:"url"`
}

func (s *Service) URL(ctx context.Context, fileID int64) (URLData, error) {
	q := url.Values{}
	q.Set("fileID", strconv.FormatInt(fileID, 10))
	return core.DoAPI[URLData](s.client, ctx, "GET", "/api/v1/direct-link/url", q, nil, true)
}

type SwitchRequest struct {
	FileID int64 `json:"fileID"`
}

type SwitchData struct {
	Filename string `json:"filename"`
}

func (s *Service) Enable(ctx context.Context, fileID int64) (SwitchData, error) {
	return core.DoAPI[SwitchData](s.client, ctx, "POST", "/api/v1/direct-link/enable", nil, SwitchRequest{FileID: fileID}, true)
}

func (s *Service) Disable(ctx context.Context, fileID int64) (SwitchData, error) {
	return core.DoAPI[SwitchData](s.client, ctx, "POST", "/api/v1/direct-link/disable", nil, SwitchRequest{FileID: fileID}, true)
}

func (s *Service) RefreshCache(ctx context.Context) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "POST", "/api/v1/direct-link/cache/refresh", nil, struct{}{}, true)
	return err
}

type DirectLinkLogItem struct {
	UniqueID      string `json:"uniqueID"`
	FileName      string `json:"fileName"`
	FileSize      int64  `json:"fileSize"`
	FilePath      string `json:"filePath"`
	DirectLinkURL string `json:"directLinkURL"`
	FileSource    int    `json:"fileSource"`
	TotalTraffic  int64  `json:"totalTraffic"`
}

type DirectLinkLogData struct {
	Total int64               `json:"total"`
	List  []DirectLinkLogItem `json:"list"`
}

type DirectLinkLogQuery struct {
	PageNum   int
	PageSize  int
	StartTime string
	EndTime   string
}

func (s *Service) Log(ctx context.Context, q DirectLinkLogQuery) (DirectLinkLogData, error) {
	qs := url.Values{}
	qs.Set("pageNum", strconv.Itoa(q.PageNum))
	qs.Set("pageSize", strconv.Itoa(q.PageSize))
	qs.Set("startTime", q.StartTime)
	qs.Set("endTime", q.EndTime)
	return core.DoAPI[DirectLinkLogData](s.client, ctx, "GET", "/api/v1/direct-link/log", qs, nil, true)
}

type DirectLinkOfflineLogItem struct {
	ID           string `json:"id"`
	FileName     string `json:"fileName"`
	FileSize     int64  `json:"fileSize"`
	LogTimeRange string `json:"logTimeRange"`
	DownloadURL  string `json:"downloadURL"`
}

type DirectLinkOfflineLogData struct {
	Total int64                      `json:"total"`
	List  []DirectLinkOfflineLogItem `json:"list"`
}

type DirectLinkOfflineLogQuery struct {
	StartHour string
	EndHour   string
	PageNum   int
	PageSize  int
}

func (s *Service) OfflineLogs(ctx context.Context, q DirectLinkOfflineLogQuery) (DirectLinkOfflineLogData, error) {
	qs := url.Values{}
	qs.Set("startHour", q.StartHour)
	qs.Set("endHour", q.EndHour)
	qs.Set("pageNum", strconv.Itoa(q.PageNum))
	qs.Set("pageSize", strconv.Itoa(q.PageSize))
	return core.DoAPI[DirectLinkOfflineLogData](s.client, ctx, "GET", "/api/v1/direct-link/offline/logs", qs, nil, true)
}

type IPBlacklistListData struct {
	IPList []string `json:"ipList"`
	Status int      `json:"status"`
}

func (s *Service) IPBlacklistList(ctx context.Context) (IPBlacklistListData, error) {
	return core.DoAPI[IPBlacklistListData](s.client, ctx, "GET", "/api/v1/developer/config/forbide-ip/list", nil, nil, true)
}

type IPBlacklistSwitchRequest struct {
	Status int `json:"Status"`
}

type IPBlacklistSwitchData struct {
	Done bool `json:"Done"`
}

func (s *Service) IPBlacklistSwitch(ctx context.Context, status int) (IPBlacklistSwitchData, error) {
	return core.DoAPI[IPBlacklistSwitchData](s.client, ctx, "POST", "/api/v1/developer/config/forbide-ip/switch", nil, IPBlacklistSwitchRequest{Status: status}, true)
}

type IPBlacklistUpdateRequest struct {
	IPList []string `json:"IpList"`
}

func (s *Service) IPBlacklistUpdate(ctx context.Context, ipList []string) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "POST", "/api/v1/developer/config/forbide-ip/update", nil, IPBlacklistUpdateRequest{IPList: ipList}, true)
	return err
}
