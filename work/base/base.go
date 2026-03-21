package base

import (
	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/work/base/account_id"
	"github.com/dysodeng/wx/work/base/contact"
)

// Base 企业微信基础模块
type Base struct {
	token   string
	aesKey  string
	account contracts.AccountInterface
}

func New(account contracts.AccountInterface, token, aesKey string) *Base {
	return &Base{
		token:   token,
		aesKey:  aesKey,
		account: account,
	}
}

// AccountId 账号ID
func (b *Base) AccountId() *account_id.AccountId {
	return account_id.New(b.account)
}

// Contact 通讯录管理
func (b *Base) Contact() *contact.Contact {
	return contact.New(b.account, b.token, b.aesKey)
}
