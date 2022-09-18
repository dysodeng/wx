package message

// Event 事件消息
type Event struct {
	Event string // 事件类型

	// 扫描带参数的二维码
	EventKey string // 事件key
	Ticket   string // 二维码ticket

	MenuId string // 菜单ID

	// 模板消息发送事件
	MsgID  int64  // 模板消息ID
	Status string // 模板消息发送状态

	// 位置上报事件
	Latitude  string
	Longitude string
	Precision string

	// 开放平台相关事件
	AppId    string
	InfoType string

	// 开放平台 component_verify_ticket 推送事件
	AuthorizerComponentVerifyTicket string

	// 开放平台授权事件
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime int64
	PreAuthCode                  string
}

// ScanEvent 扫描带参数的二维码事件
type ScanEvent struct {
	EventKey string // 事件key
	Ticket   string // 二维码ticket
}

// LocationEvent 上报地理位置事件
type LocationEvent struct {
	Latitude  string
	Longitude string
	Precision string
}

// MenuEvent 菜单事件
type MenuEvent struct {
	EventKey string // 事件key
	MenuId   string // 菜单ID
}

// TemplateSendJobFinishEvent 模板消息推荐通知事件
type TemplateSendJobFinishEvent struct {
	MsgID  int64  // 模板消息ID
	Status string // 模板消息发送状态
}

// ComponentVerifyTicketEvent 开放平台 component_verify_ticket 推送事件
type ComponentVerifyTicketEvent struct {
	AppId                 string
	InfoType              string
	ComponentVerifyTicket string
}

// AuthorizerEvent 开放平台公众账号授权事件
type AuthorizerEvent struct {
	AppId                        string
	InfoType                     string
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime int64
	PreAuthCode                  string
}

// Scan 扫描带参数的二维码事件
func (e *Event) Scan() *ScanEvent {
	return &ScanEvent{
		EventKey: e.EventKey,
		Ticket:   e.Ticket,
	}
}

// Location 位置上报事件
func (e *Event) Location() *LocationEvent {
	return &LocationEvent{
		Latitude:  e.Latitude,
		Longitude: e.Longitude,
		Precision: e.Precision,
	}
}

// Menu 菜单事件
func (e *Event) Menu() *MenuEvent {
	return &MenuEvent{
		EventKey: e.EventKey,
		MenuId:   e.MenuId,
	}
}

// TemplateSendJobFinish 模板消息推荐状态事件
func (e *Event) TemplateSendJobFinish() *TemplateSendJobFinishEvent {
	return &TemplateSendJobFinishEvent{
		MsgID:  e.MsgID,
		Status: e.Status,
	}
}

// ComponentVerifyTicket 开放平台 component_verify_ticket 推送事件
func (e *Event) ComponentVerifyTicket() *ComponentVerifyTicketEvent {
	return &ComponentVerifyTicketEvent{
		AppId:                 e.AppId,
		InfoType:              e.InfoType,
		ComponentVerifyTicket: e.AuthorizerComponentVerifyTicket,
	}
}

// Authorizer 开放平台-公众账号授权事件
func (e *Event) Authorizer() *AuthorizerEvent {
	return &AuthorizerEvent{
		AppId:                        e.AppId,
		InfoType:                     e.InfoType,
		AuthorizerAppid:              e.AuthorizerAppid,
		AuthorizationCode:            e.AuthorizationCode,
		AuthorizationCodeExpiredTime: e.AuthorizationCodeExpiredTime,
		PreAuthCode:                  e.PreAuthCode,
	}
}
