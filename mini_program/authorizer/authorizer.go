package authorizer

import "github.com/dysodeng/wx/kernel/contracts"

// Authorizer 小程序授权开放平台后的相关接口
// 此类接口只能由授权到开放平台的小程序调用或直接由开放平台调用
type Authorizer struct {
	account contracts.AccountInterface
}

func NewAuthorizer(account contracts.AccountInterface) *Authorizer {
	return &Authorizer{account: account}
}

// Account 小程序基础信息
func (authorizer *Authorizer) Account() *Account {
	return NewAccount(authorizer.account)
}

// Domain 小程序域名配置
func (authorizer *Authorizer) Domain() *Domain {
	return NewDomain(authorizer.account)
}
