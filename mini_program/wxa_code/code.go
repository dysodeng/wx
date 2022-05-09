package wxa_code

import "github.com/dysodeng/wx/kernel/contracts"

// WxaCode 小程序码
type WxaCode struct {
	account contracts.AccountInterface
}

func NewWxaCode(account contracts.AccountInterface) *WxaCode {
	return &WxaCode{account: account}
}
