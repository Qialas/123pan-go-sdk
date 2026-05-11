package sharemgmt

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

type CreateShareRequest struct {
	ShareName          string `json:"shareName"`
	ShareExpire        int    `json:"shareExpire"`
	FileIDList         string `json:"fileIDList"`
	SharePwd           string `json:"sharePwd,omitempty"`
	TrafficSwitch      *int   `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch *int   `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       *int64 `json:"trafficLimit,omitempty"`
}

type CreateShareData struct {
	ShareID  int64  `json:"shareID"`
	ShareKey string `json:"shareKey"`
}

func (s *Service) Create(ctx context.Context, req CreateShareRequest) (CreateShareData, error) {
	return core.DoAPI[CreateShareData](s.client, ctx, "POST", "/api/v1/share/create", nil, req, true)
}

type ShareListItem struct {
	ShareID            int64  `json:"shareId"`
	ShareKey           string `json:"shareKey"`
	ShareName          string `json:"shareName"`
	Expiration         string `json:"expiration"`
	Expired            int    `json:"expired"`
	SharePwd           string `json:"sharePwd"`
	TrafficSwitch      int    `json:"trafficSwitch"`
	TrafficLimitSwitch int    `json:"trafficLimitSwitch"`
	TrafficLimit       int64  `json:"trafficLimit"`
	BytesCharge        int64  `json:"bytesCharge"`
	PreviewCount       int64  `json:"previewCount"`
	DownloadCount      int64  `json:"downloadCount"`
	SaveCount          int64  `json:"saveCount"`
}

type ShareListData struct {
	LastShareID int64           `json:"lastShareId"`
	ShareList   []ShareListItem `json:"shareList"`
}

func (s *Service) List(ctx context.Context, limit int, lastShareID *int64) (ShareListData, error) {
	q := url.Values{}
	q.Set("limit", strconv.Itoa(limit))
	if lastShareID != nil {
		q.Set("lastShareId", strconv.FormatInt(*lastShareID, 10))
	}
	return core.DoAPI[ShareListData](s.client, ctx, "GET", "/api/v1/share/list", q, nil, true)
}

type UpdateShareRequest struct {
	ShareIDList        []uint64 `json:"shareIdList"`
	TrafficSwitch      *int     `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch *int     `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       *int64   `json:"trafficLimit,omitempty"`
}

func (s *Service) Update(ctx context.Context, req UpdateShareRequest) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "PUT", "/api/v1/share/list/info", nil, req, true)
	return err
}

type CreatePaidShareRequest struct {
	ShareName          string `json:"shareName"`
	FileIDList         string `json:"fileIDList"`
	PayAmount          int    `json:"payAmount"`
	IsReward           *int   `json:"isReward,omitempty"`
	ResourceDesc       string `json:"resourceDesc,omitempty"`
	TrafficSwitch      *int   `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch *int   `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       *int64 `json:"trafficLimit,omitempty"`
}

type PaidShareListItem struct {
	ShareID            int64  `json:"shareId"`
	ShareKey           string `json:"shareKey"`
	ShareName          string `json:"shareName"`
	PayAmount          int64  `json:"payAmount"`
	Amount             int64  `json:"amount"`
	Expiration         string `json:"expiration"`
	Expired            int    `json:"expired"`
	TrafficSwitch      int    `json:"trafficSwitch"`
	TrafficLimitSwitch int    `json:"trafficLimitSwitch"`
	TrafficLimit       int64  `json:"trafficLimit"`
	BytesCharge        int64  `json:"bytesCharge"`
	PreviewCount       int64  `json:"previewCount"`
	DownloadCount      int64  `json:"downloadCount"`
	SaveCount          int64  `json:"saveCount"`
}

type PaidShareListData struct {
	LastShareID int64               `json:"lastShareId"`
	ShareList   []PaidShareListItem `json:"shareList"`
}

func (s *Service) CreatePaid(ctx context.Context, req CreatePaidShareRequest) (CreateShareData, error) {
	return core.DoAPI[CreateShareData](s.client, ctx, "POST", "/api/v1/share/content-payment/create", nil, req, true)
}

func (s *Service) ListPaid(ctx context.Context, limit int, lastShareID *int64) (PaidShareListData, error) {
	q := url.Values{}
	q.Set("limit", strconv.Itoa(limit))
	if lastShareID != nil {
		q.Set("lastShareId", strconv.FormatInt(*lastShareID, 10))
	}
	return core.DoAPI[PaidShareListData](s.client, ctx, "GET", "/api/v1/share/payment/list", q, nil, true)
}

type UpdatePaidShareRequest struct {
	ShareIDList        []uint64 `json:"shareIdList"`
	TrafficSwitch      *int     `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch *int     `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       *int64   `json:"trafficLimit,omitempty"`
}

func (s *Service) UpdatePaid(ctx context.Context, req UpdatePaidShareRequest) error {
	_, err := core.DoAPI[struct{}](s.client, ctx, "PUT", "/api/v1/share/payment/list/info", nil, req, true)
	return err
}
