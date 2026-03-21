package message

import "encoding/xml"

// OpenPlatformEvent 开放平台事件消息
type OpenPlatformEvent struct {
	AppId    string
	InfoType string

	// component_verify_ticket 推送事件
	ComponentVerifyTicket string

	// 授权事件
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime int64
	PreAuthCode                  string

	Ret      int    `xml:"ret"`
	Nickname string `xml:"nickname"`
	Reason   string `xml:"reason"`
	First    int64  `xml:"first"`
	Second   int64  `xml:"second"`

	SuccTime   int64
	FailTime   int64
	DelayTime  int64
	ScreenShot string
}

// parseOpenPlatformEvent 从 RawBody 解析开放平台事件
func parseOpenPlatformEvent(rawBody []byte) *OpenPlatformEvent {
	if len(rawBody) == 0 {
		return &OpenPlatformEvent{}
	}
	var op struct {
		XMLName                      xml.Name `xml:"xml"`
		AppId                        string
		InfoType                     string
		ComponentVerifyTicket        string
		AuthorizerAppid              string
		AuthorizationCode            string
		AuthorizationCodeExpiredTime int64
		PreAuthCode                  string
		Ret                          int    `xml:"ret"`
		Nickname                     string `xml:"nickname"`
		Reason                       string `xml:"reason"`
		First                        int64  `xml:"first"`
		Second                       int64  `xml:"second"`
		SuccTime                     int64
		FailTime                     int64
		DelayTime                    int64
		ScreenShot                   string
	}
	if xml.Unmarshal(rawBody, &op) != nil {
		return &OpenPlatformEvent{}
	}
	return &OpenPlatformEvent{
		AppId:                        op.AppId,
		InfoType:                     op.InfoType,
		ComponentVerifyTicket:        op.ComponentVerifyTicket,
		AuthorizerAppid:              op.AuthorizerAppid,
		AuthorizationCode:            op.AuthorizationCode,
		AuthorizationCodeExpiredTime: op.AuthorizationCodeExpiredTime,
		PreAuthCode:                  op.PreAuthCode,
		Ret:                          op.Ret,
		Nickname:                     op.Nickname,
		Reason:                       op.Reason,
		First:                        op.First,
		Second:                       op.Second,
		SuccTime:                     op.SuccTime,
		FailTime:                     op.FailTime,
		DelayTime:                    op.DelayTime,
		ScreenShot:                   op.ScreenShot,
	}
}
