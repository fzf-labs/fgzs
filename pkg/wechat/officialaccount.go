package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"sync"
)

type OfficialAccountConfig struct {
	AppID     string `json:"AppId"`     //appid
	AppSecret string `json:"AppSecret"` //appsecret
}

var (
	officialAccountOnce sync.Once
	WxOfficialAccount   OfficialAccount
)

type OfficialAccount struct {
	config          *OfficialAccountConfig
	officialAccount *officialaccount.OfficialAccount
}

func NewOfficialAccount(cfg *OfficialAccountConfig, cache *RedisCache) *OfficialAccount {
	officialAccountOnce.Do(
		func() {
			account := wechat.NewWechat().GetOfficialAccount(&config.Config{
				AppID:     cfg.AppID,
				AppSecret: cfg.AppSecret,
				Token:     "",
				//EncodingAESKey: "",
				Cache: cache,
			})
			WxOfficialAccount.config = cfg
			WxOfficialAccount.officialAccount = account
		})
	return &WxOfficialAccount
}

// GetWxUserByCode 根据前端返回的code获取用户openid 或者 unionid
func (oa *OfficialAccount) GetWxUserByCode(code string) (oauth.ResAccessToken, error) {
	return oa.officialAccount.GetOauth().GetUserAccessToken(code)
}

// GetUserInfo 获取用户信息
func (oa *OfficialAccount) GetUserInfo(accessToken, openID string) (oauth.UserInfo, error) {
	return oa.officialAccount.GetOauth().GetUserInfo(accessToken, openID, "")
}
