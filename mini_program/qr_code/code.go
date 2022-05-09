package qr_code

import "github.com/dysodeng/wx/kernel/contracts"

// QrCode 普通链接二维码
type QrCode struct {
	account contracts.AccountInterface
}

func NewQrCode(account contracts.AccountInterface) *QrCode {
	return &QrCode{account: account}
}
