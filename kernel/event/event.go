package event

// EventType 事件处理类型
type EventType string

const (
	All EventType = "*"

	// 公众号事件消息
	Event       EventType = "event"
	Subscribe   EventType = "subscribe"
	Unsubscribe EventType = "unsubscribe"
	Scan        EventType = "scan"
	Location    EventType = "location"

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

	// 企业微信通讯录变更事件
	ChangeContact EventType = "change_contact"

	// 企业微信异步任务完成事件
	BatchJobResult EventType = "batch_job_result"

	// 企业微信应用事件
	EnterAgent            EventType = "enter_agent"              // 进入应用
	ChangeExternalChat    EventType = "change_external_chat"     // 客户群变更事件
	ChangeExternalContact EventType = "change_external_contact"  // 外部联系人变更事件
	ChangeExternalTag     EventType = "change_external_tag"      // 企业客户标签变更事件
	SysApprovalChange     EventType = "sys_approval_change"      // 审批状态变更事件
	OpenApprovalChange    EventType = "open_approval_change"     // 自建审批申请状态变化回调
	ShareAgentChange      EventType = "share_agent_change"       // 共享应用事件回调
	ShareChain            EventType = "share_chain"              // 上下游共享应用事件回调
	TemplateCardEvent     EventType = "template_card_event"      // 模板卡片事件推送
	TemplateCardMenu      EventType = "template_card_menu_event" // 通用模板卡片右上角菜单事件推送
	LivingStatusChange    EventType = "living_status_change"     // 直播事件回调
	MsgAuditNotify        EventType = "msgaudit_notify"          // 会话存档事件回调
)
