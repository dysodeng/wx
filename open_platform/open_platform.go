package open_platform

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/official"
)

// OpenPlatform 开放平台
type OpenPlatform struct {
	config *config
	option *option
}

func NewOpenPlatform() *OpenPlatform {
	return &OpenPlatform{}
}

// Server 服务端
func (open *OpenPlatform) Server() *base.Server {
	return base.NewServer(open)
}

// Official 授权到开放平台的公众号
// @param appId string 公众号appID
// @param authorizerRefreshToken string 公众号授权刷新token
func (open *OpenPlatform) Official(appId, authorizerRefreshToken string) *official.Official {
	return official.NewOfficialWithOpenPlatform(appId, authorizerRefreshToken, open)
}
