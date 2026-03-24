package kf

import "github.com/dysodeng/wx/kernel/contracts"

// Kf 微信客服
type Kf struct {
	account contracts.AccountInterface
}

func NewKf(account contracts.AccountInterface) *Kf {
	return &Kf{account: account}
}

// Account 客服账号管理
func (k *Kf) Account() *Account {
	return NewAccount(k.account)
}
