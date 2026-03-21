package event

// EventType 事件处理类型
type EventType string

const (
	All EventType = "*"

	// 公众号事件消息
	Event     EventType = "event"
	Subscribe EventType = "subscribe"
	Scan      EventType = "scan"
	Location  EventType = "location"

	// 公众号菜单事件消息
	Click           EventType = "click"
	View            EventType = "view"
	ScancodePush    EventType = "scancode_push"
	ScancodeWaitMsg EventType = "scancode_waitmsg"
	PicSysPhoto     EventType = "pic_sysphoto"
	PicPhotoOrAlbum EventType = "pic_photo_or_album"
	PicWeiXin       EventType = "pic_weixin"
	LocationSelect  EventType = "location_select"
	ViewMiniProgram EventType = "view_miniprogram"

	// 公众号对话能力事件消息
	GuideQrcodeScanEvent EventType = "guide_qrcode_scan_event"

	// 公众号群发事件消息
	MassSendJobFinish EventType = "masssendjobfinish"
	PublishJobFinish  EventType = "publishjobfinish"

	// 微信认证通知事件消息
	QualificationVerifySuccess EventType = "qualification_verify_success"
	QualificationVerifyFail    EventType = "qualification_verify_fail"
	NamingVerifySuccess        EventType = "naming_verify_success"
	NamingVerifyFail           EventType = "naming_verify_fail"
	AnnualRenew                EventType = "annual_renew"
	VerifyExpired              EventType = "verify_expired"

	// 微信卡券事件消息
	CardPassCheck            EventType = "card_pass_check"
	CardNotPassCheck         EventType = "card_not_pass_check"
	UserGetCard              EventType = "user_get_card"
	UserGiftingCard          EventType = "user_gifting_card"
	UserDelCard              EventType = "user_del_card"
	UserConsumeCard          EventType = "user_consume_card"
	UserPayFromPayCell       EventType = "user_pay_from_pay_cell"
	UserViewCard             EventType = "user_view_card"
	UserEnterSessionFromCard EventType = "user_enter_session_from_card"
	UpdateMemberCard         EventType = "update_member_card"
	CardSkuRemind            EventType = "card_sku_remind"
	CardPayOrder             EventType = "card_pay_order"
	SubmitMemberCardUserInfo EventType = "submit_membercard_user_info"

	// 开放平台事件
	Authorized               EventType = "authorized"
	Unauthorized             EventType = "unauthorized"
	UpdateAuthorized         EventType = "updateauthorized"
	ComponentVerifyTicket    EventType = "component_verify_ticket"
	NotifyThirdFasteRegister EventType = "notify_third_fasteregister"
	WxaNicknameAudit         EventType = "wxa_nickname_audit"
	WeAppAuditSuccess        EventType = "weapp_audit_success"
	WeAppAuditFail           EventType = "weapp_audit_fail"
	WeAppAuditDelay          EventType = "weapp_audit_delay"
)
