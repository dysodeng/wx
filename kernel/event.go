package kernel

// EventGuard 服务类型
type EventGuard string

const (
	GuardAll EventGuard = "*"

	/* 公众号普通消息 */

	/* 公众号事件消息 */
	GuardEvent          EventGuard = "event"
	GuardEventSubscribe EventGuard = "subscribe"
	GuardEventScan      EventGuard = "scan"
	GuardEventLocation  EventGuard = "location"

	/* 公众号菜单事件消息 */
	GuardEventClick           EventGuard = "click"
	GuardEventView            EventGuard = "view"
	GuardEventScancodePush    EventGuard = "scancode_push"
	GuardEventScancodeWaitMsg EventGuard = "scancode_waitmsg"
	GuardEventPicSysPhoto     EventGuard = "pic_sysphoto"
	GuardEventPicPhotoOrAlbum EventGuard = "pic_photo_or_album"
	GuardEventPicWeiXin       EventGuard = "pic_weixin"
	GuardEventLocationSelect  EventGuard = "location_select"
	GuardEventViewMiniProgram EventGuard = "view_miniprogram"

	/* 公众号对话能力事件消息 */
	GuardEventGuideQrcodeScanEvent EventGuard = "guide_qrcode_scan_event"

	/* 公众号群发事件消息 */
	GuardEventMassSendJobFinish EventGuard = "masssendjobfinish"
	GuardEventPublishJobFinish  EventGuard = "publishjobfinish"

	/* 微信认证通知事件消息 */
	GuardEventQualificationVerifySuccess EventGuard = "qualification_verify_success"
	GuardEventQualificationVerifyFail    EventGuard = "qualification_verify_fail"
	GuardEventNamingVerifySuccess        EventGuard = "naming_verify_success"
	GuardEventNamingVerifyFail           EventGuard = "naming_verify_fail"
	GuardEventAnnualRenew                EventGuard = "annual_renew"
	GuardEventVerifyExpired              EventGuard = "verify_expired"

	/* 微信卡券事件消息 */
	GuardEventCardPassCheck            EventGuard = "card_pass_check"
	GuardEventCardNotPassCheck         EventGuard = "card_not_pass_check"
	GuardEventUserGetCard              EventGuard = "user_get_card"
	GuardEventUserGiftingCard          EventGuard = "user_gifting_card"
	GuardEventUserDelCard              EventGuard = "user_del_card"
	GuardEventUserConsumeCard          EventGuard = "user_consume_card"
	GuardEventUserPayFromPayCell       EventGuard = "user_pay_from_pay_cell"
	GuardEventUserViewCard             EventGuard = "user_view_card"
	GuardEventUserEnterSessionFromCard EventGuard = "user_enter_session_from_card"
	GuardEventUpdateMemberCard         EventGuard = "update_member_card"
	GuardEventCardSkuRemind            EventGuard = "card_sku_remind"
	GuardEventCardPayOrder             EventGuard = "card_pay_order"
	GuardEventSubmitMemberCardUserInfo EventGuard = "submit_membercard_user_info"

	/* 开放平台事件 */
	GuardEventAuthorized               EventGuard = "authorized"
	GuardEventUnauthorized             EventGuard = "unauthorized"
	GuardEventUpdateAuthorized         EventGuard = "updateauthorized"
	GuardEventComponentVerifyTicket    EventGuard = "component_verify_ticket"
	GuardEventNotifyThirdFasteRegister EventGuard = "notify_third_fasteregister"
)
