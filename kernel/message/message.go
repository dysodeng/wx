package message

import (
	"encoding/xml"
	"time"
)

// Message 消息体
type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MsgId        int

	// 文本消息
	Content string

	// 媒体消息通用字段
	MediaId string

	// 图片消息
	PicUrl string

	// 语音消息
	Format      string
	Recognition string

	// 视频消息
	ThumbMediaId string

	// 地理位置消息
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string
	Label     string

	// 链接消息
	Title       string
	Description string
	Url         string

	// 事件消息
	Event    string // 事件类型
	EventKey string // 事件key
	Ticket   string // 二维码ticket

	// 模板消息发送事件
	MsgID  string // 模板消息ID
	Status int    // 模板消息发送状态

	// 位置上报事件
	Latitude  string
	Longitude string
	Precision string

	// 开放平台
	AppId                 string
	InfoType              string
	ComponentVerifyTicket string

	// 开放平台授权事件
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime int64
	PreAuthCode                  string

	Ret      int    `xml:"ret"`
	Nickname string `xml:"nickname"`
	Reason   string `xml:"reason"`
	First    int64  `xml:"first"`
	Second   int64  `xml:"second"`
}
