package authorizer

import (
	"github.com/dysodeng/wx/base/open"
	"github.com/dysodeng/wx/kernel/contracts"
)

// Authorizer 小程序授权开放平台后的相关接口
// 此类接口只能由授权到开放平台的小程序调用或直接由开放平台调用
type Authorizer struct {
	account contracts.AccountInterface
}

func New(account contracts.AccountInterface) *Authorizer {
	return &Authorizer{account: account}
}

// Open 开放平台
func (authorizer *Authorizer) Open() *open.Open {
	return open.New(authorizer.account)
}
