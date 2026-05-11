package usermgmt

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

type AccessTokenRequest struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

type AccessTokenData struct {
	AccessToken string `json:"accessToken"`
	ExpiredAt   string `json:"expiredAt"`
}

func (s *Service) GetAccessToken(ctx context.Context, clientID, clientSecret string) (AccessTokenData, error) {
	req := AccessTokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	return core.DoAPI[AccessTokenData](s.client, ctx, "POST", "/api/v1/access_token", url.Values{}, req, false)
}

type OAuthAccessTokenData struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
}

type OAuthAccessTokenQuery struct {
	ClientID     string
	ClientSecret string
	GrantType    string
	Code         string
	RefreshToken string
	RedirectURI  string
}

func (s *Service) OAuthAccessToken(ctx context.Context, q OAuthAccessTokenQuery) (OAuthAccessTokenData, error) {
	qs := url.Values{}
	qs.Set("client_id", q.ClientID)
	qs.Set("client_secret", q.ClientSecret)
	qs.Set("grant_type", q.GrantType)
	if q.Code != "" {
		qs.Set("code", q.Code)
	}
	if q.RefreshToken != "" {
		qs.Set("refresh_token", q.RefreshToken)
	}
	if q.RedirectURI != "" {
		qs.Set("redirect_uri", q.RedirectURI)
	}
	return core.DoAPI[OAuthAccessTokenData](s.client, ctx, "POST", "/api/v1/oauth2/access_token", qs, nil, false)
}

type UserInfoData struct {
	UID            int64          `json:"uid"`
	Nickname       string         `json:"nickname"`
	HeadImage      string         `json:"headImage"`
	Passport       string         `json:"passport"`
	Mail           string         `json:"mail"`
	SpaceUsed      int64          `json:"spaceUsed"`
	SpacePermanent int64          `json:"spacePermanent"`
	SpaceTemp      int64          `json:"spaceTemp"`
	SpaceTempExpr  string         `json:"spaceTempExpr"`
	VIP            bool           `json:"vip"`
	DirectTraffic  int64          `json:"directTraffic"`
	IsHideUID      bool           `json:"isHideUID"`
	HTTPSCount     int64          `json:"httpsCount"`
	VIPInfo        *VIPInfo       `json:"vipInfo"`
	DeveloperInfo  *DeveloperInfo `json:"developerInfo"`
}

type VIPInfo struct {
	VIPLevel  int    `json:"vipLevel"`
	VIPLabel  string `json:"vipLabel"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type DeveloperInfo struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (s *Service) Info(ctx context.Context) (UserInfoData, error) {
	return core.DoAPI[UserInfoData](s.client, ctx, "GET", "/api/v1/user/info", nil, nil, true)
}

type ReferralBonusListItem struct {
	ID          int64  `json:"id"`
	UID         int64  `json:"uid"`
	Mobile      string `json:"mobile"`
	SourceType  int    `json:"sourceType"`
	Paid        int64  `json:"paid"`
	Bonus       int64  `json:"bonus"`
	ProductName string `json:"productName"`
	CreateTime  string `json:"createTime"`
}

type ReferralBonusListData struct {
	LastID int64                   `json:"lastId"`
	List   []ReferralBonusListItem `json:"list"`
}

type ReferralBonusListQuery struct {
	Limit       int
	TimeStart   string
	TimeEnd     string
	LastID      *int64
	UID         *int64
	BonusStatus *int
}

func (s *Service) ReferralBonusList(ctx context.Context, q ReferralBonusListQuery) (ReferralBonusListData, error) {
	qs := url.Values{}
	qs.Set("limit", strconv.Itoa(q.Limit))
	qs.Set("timeStart", q.TimeStart)
	qs.Set("timeEnd", q.TimeEnd)
	if q.LastID != nil {
		qs.Set("lastId", strconv.FormatInt(*q.LastID, 10))
	}
	if q.UID != nil {
		qs.Set("uid", strconv.FormatInt(*q.UID, 10))
	}
	if q.BonusStatus != nil {
		qs.Set("bonusStatus", strconv.Itoa(*q.BonusStatus))
	}
	return core.DoAPI[ReferralBonusListData](s.client, ctx, "GET", "/api/v1/referral/bonus/list", qs, nil, true)
}
