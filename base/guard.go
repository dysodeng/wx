package base

// Guard 服务类型
type Guard string

const (
	GuardAll Guard = "*"

	/* 公众号普通消息 */

	/* 公众号事件消息 */
	GuardEvent          Guard = "event"
	GuardEventSubscribe Guard = "subscribe"
	GuardEventScan      Guard = "scan"
	GuardEventLocation  Guard = "location"

	/* 公众号菜单事件消息 */
	GuardEventClick           Guard = "click"
	GuardEventView            Guard = "view"
	GuardEventScancodePush    Guard = "scancode_push"
	GuardEventScancodeWaitMsg Guard = "scancode_waitmsg"
	GuardEventPicSysPhoto     Guard = "pic_sysphoto"
	GuardEventPicPhotoOrAlbum Guard = "pic_photo_or_album"
	GuardEventPicWeixin       Guard = "pic_weixin"
	GuardEventLocationSelect  Guard = "location_select"
	GuardEventViewMiniProgram Guard = "view_miniprogram"

	/* 公众号对话能力事件消息 */
	GuardEventGuideQrcodeScanEvent Guard = "guide_qrcode_scan_event"

	/* 公众号群发事件消息 */
	GuardEventMassSendJobFinish Guard = "masssendjobfinish"
	GuardEventPublishJobFinish  Guard = "publishjobfinish"

	/* 微信认证通知事件消息 */
	GuardEventQualificationVerifySuccess Guard = "qualification_verify_success"
	GuardEventQualificationVerifyFail    Guard = "qualification_verify_fail"
	GuardEventNamingVerifySuccess        Guard = "naming_verify_success"
	GuardEventNamingVerifyFail           Guard = "naming_verify_fail"
	GuardEventAnnualRenew                Guard = "annual_renew"
	GuardEventVerifyExpired              Guard = "verify_expired"

	/* 微信卡券事件消息 */
	GuardEventCardPassCheck            Guard = "card_pass_check"
	GuardEventCardNotPassCheck         Guard = "card_not_pass_check"
	GuardEventUserGetCard              Guard = "user_get_card"
	GuardEventUserGiftingCard          Guard = "user_gifting_card"
	GuardEventUserDelCard              Guard = "user_del_card"
	GuardEventUserConsumeCard          Guard = "user_consume_card"
	GuardEventUserPayFromPayCell       Guard = "user_pay_from_pay_cell"
	GuardEventUserViewCard             Guard = "user_view_card"
	GuardEventUserEnterSessionFromCard Guard = "user_enter_session_from_card"
	GuardEventUpdateMemberCard         Guard = "update_member_card"
	GuardEventCardSkuRemind            Guard = "card_sku_remind"
	GuardEventCardPayOrder             Guard = "card_pay_order"
	GuardEventSubmitMemberCardUserInfo Guard = "submit_membercard_user_info"
)
