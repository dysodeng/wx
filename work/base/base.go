package base

import (
	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/work/base/account_id"
	"github.com/dysodeng/wx/work/base/contact"
	"github.com/dysodeng/wx/work/base/oauth"
)

// Base 企业微信基础模块
type Base struct {
	account contracts.AccountInterface
}

func New(account contracts.AccountInterface) *Base {
	return &Base{account: account}
}

// AccountId 账号ID
func (b *Base) AccountId() *account_id.AccountId {
	return account_id.New(b.account)
}

// Contact 通讯录管理
func (b *Base) Contact() *contact.Contact {
	return contact.New(b.account)
}

// OAuth 身份验证-网页授权登录
func (b *Base) OAuth() *oauth.OAuth {
	return oauth.New(b.account)
}
