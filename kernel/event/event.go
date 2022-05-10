package event

// Guard 事件类型
type Guard string

const (
	All Guard = "*"

	/* 公众号普通消息 */

	/* 公众号事件消息 */
	Event     Guard = "event"
	Subscribe Guard = "subscribe"
	Scan      Guard = "scan"
	Location  Guard = "location"

	/* 公众号菜单事件消息 */
	Click           Guard = "click"
	View            Guard = "view"
	ScancodePush    Guard = "scancode_push"
	ScancodeWaitMsg Guard = "scancode_waitmsg"
	PicSysPhoto     Guard = "pic_sysphoto"
	PicPhotoOrAlbum Guard = "pic_photo_or_album"
	PicWeiXin       Guard = "pic_weixin"
	LocationSelect  Guard = "location_select"
	ViewMiniProgram Guard = "view_miniprogram"

	/* 公众号对话能力事件消息 */
	GuideQrcodeScanEvent Guard = "guide_qrcode_scan_event"

	/* 公众号群发事件消息 */
	MassSendJobFinish Guard = "masssendjobfinish"
	PublishJobFinish  Guard = "publishjobfinish"

	/* 微信认证通知事件消息 */
	QualificationVerifySuccess Guard = "qualification_verify_success"
	QualificationVerifyFail    Guard = "qualification_verify_fail"
	NamingVerifySuccess        Guard = "naming_verify_success"
	NamingVerifyFail           Guard = "naming_verify_fail"
	AnnualRenew                Guard = "annual_renew"
	VerifyExpired              Guard = "verify_expired"

	/* 微信卡券事件消息 */
	CardPassCheck            Guard = "card_pass_check"
	CardNotPassCheck         Guard = "card_not_pass_check"
	UserGetCard              Guard = "user_get_card"
	UserGiftingCard          Guard = "user_gifting_card"
	UserDelCard              Guard = "user_del_card"
	UserConsumeCard          Guard = "user_consume_card"
	UserPayFromPayCell       Guard = "user_pay_from_pay_cell"
	UserViewCard             Guard = "user_view_card"
	UserEnterSessionFromCard Guard = "user_enter_session_from_card"
	UpdateMemberCard         Guard = "update_member_card"
	CardSkuRemind            Guard = "card_sku_remind"
	CardPayOrder             Guard = "card_pay_order"
	SubmitMemberCardUserInfo Guard = "submit_membercard_user_info"

	/* 开放平台事件 */
	Authorized               Guard = "authorized"
	Unauthorized             Guard = "unauthorized"
	UpdateAuthorized         Guard = "updateauthorized"
	ComponentVerifyTicket    Guard = "component_verify_ticket"
	NotifyThirdFasteRegister Guard = "notify_third_fasteregister"
	WxaNicknameAudit         Guard = "wxa_nickname_audit"
)
