package message

import (
	"encoding/xml"
	"time"
)

// Message 消息体
type Message struct {
	// 原始 XML 数据，供平台特定解析使用
	RawBody []byte `xml:"-"`

	// 消息头
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MsgId        int64
	AgentID      string // 企业微信应用ID

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
	Event string // 事件类型

	EventKey string // 事件key (扫描带参数的二维码、菜单点击事件)
	Ticket   string // 二维码ticket

	MenuId string // 菜单ID

	// 开放平台消息路由字段
	InfoType string

	// 企业微信通讯录变更事件
	ChangeType string

	// 位置上报事件
	Latitude  string
	Longitude string
	Precision string

	// 模板消息发送事件
	MsgID  int64  // 模板消息ID
	Status string // 模板消息发送状态
}

// Header 消息头
type Header struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MsgId        int64
	AgentID      string
}

// Text 文本消息
type Text struct {
	Content string
}

// Image 图片消息
type Image struct {
	MediaId string
	PicUrl  string
}

// Voice 语音消息
type Voice struct {
	MediaId     string
	Format      string
	Recognition string
}

// Video 视频消息
type Video struct {
	MediaId      string
	ThumbMediaId string
}

// Location 地理位置消息
type Location struct {
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string
	Label     string
}

// Link 链接消息
type Link struct {
	Title       string
	Description string
	Url         string
}

// Header 消息头
func (m *Message) Header() *Header {
	return &Header{
		ToUserName:   m.ToUserName,
		FromUserName: m.FromUserName,
		CreateTime:   m.CreateTime,
		MsgType:      m.MsgType,
		MsgId:        m.MsgId,
		AgentID:      m.AgentID,
	}
}

// Text 文本消息
func (m *Message) Text() *Text {
	return &Text{
		Content: m.Content,
	}
}

// Image 图片消息
func (m *Message) Image() *Image {
	return &Image{
		MediaId: m.MediaId,
		PicUrl:  m.PicUrl,
	}
}

// Voice 语音消息
func (m *Message) Voice() *Voice {
	return &Voice{
		MediaId:     m.MediaId,
		Format:      m.Format,
		Recognition: m.Recognition,
	}
}

// Video 视频消息
func (m *Message) Video() *Video {
	return &Video{
		MediaId:      m.MediaId,
		ThumbMediaId: m.ThumbMediaId,
	}
}

// ShortVideo 短视频消息
func (m *Message) ShortVideo() *Video {
	return &Video{
		MediaId:      m.MediaId,
		ThumbMediaId: m.ThumbMediaId,
	}
}

// Location 位置消息
func (m *Message) Location() *Location {
	return &Location{
		LocationX: m.LocationX,
		LocationY: m.LocationY,
		Scale:     m.Scale,
		Label:     m.Label,
	}
}

// Link 链接消息
func (m *Message) Link() *Link {
	return &Link{
		Title:       m.Title,
		Description: m.Description,
		Url:         m.Url,
	}
}

// EventMessage 事件消息
func (m *Message) EventMessage() *Event {
	return &Event{
		Event:     m.Event,
		EventKey:  m.EventKey,
		Ticket:    m.Ticket,
		MsgID:     m.MsgID,
		Status:    m.Status,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
		Precision: m.Precision,
	}
}

// OpenPlatformEvent 开放平台事件消息
func (m *Message) OpenPlatformEvent() *OpenPlatformEvent {
	return parseOpenPlatformEvent(m.RawBody)
}

// WorkContactEvent 企业微信通讯录变更事件消息
func (m *Message) WorkContactEvent() *WorkContactEvent {
	return parseWorkContactEvent(m.RawBody)
}

// WorkBatchJobResultEvent 企业微信异步任务完成事件消息
func (m *Message) WorkBatchJobResultEvent() *WorkBatchJobResultEvent {
	return parseWorkBatchJobResultEvent(m.RawBody)
}

// WorkExternalContactEvent 企业微信外部联系人变更事件消息
func (m *Message) WorkExternalContactEvent() *WorkExternalContactEvent {
	return parseWorkExternalContactEvent(m.RawBody)
}

// WorkExternalChatEvent 企业微信客户群变更事件消息
func (m *Message) WorkExternalChatEvent() *WorkExternalChatEvent {
	return parseWorkExternalChatEvent(m.RawBody)
}

// WorkExternalTagEvent 企业微信企业客户标签变更事件消息
func (m *Message) WorkExternalTagEvent() *WorkExternalTagEvent {
	return parseWorkExternalTagEvent(m.RawBody)
}

// WorkTemplateCardEvent 企业微信模板卡片事件消息
func (m *Message) WorkTemplateCardEvent() *WorkTemplateCardEvent {
	return parseWorkTemplateCardEvent(m.RawBody)
}

// WorkLivingStatusChangeEvent 企业微信直播事件消息
func (m *Message) WorkLivingStatusChangeEvent() *WorkLivingStatusChangeEvent {
	return parseWorkLivingStatusChangeEvent(m.RawBody)
}

// WorkApprovalEvent 企业微信审批事件消息
func (m *Message) WorkApprovalEvent() *WorkApprovalEvent {
	return parseWorkApprovalEvent(m.RawBody)
}
