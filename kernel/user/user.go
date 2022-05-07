package user

import (
	baseError "github.com/dysodeng/wx/kernel/error"
)

type User struct {
	baseError.WxApiError
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        uint8    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionid"`
}
