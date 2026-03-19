package authorizer

import "github.com/dysodeng/wx/kernel/contracts"

// Authorizer 企业微信公众账号授权
type Authorizer struct {
	account contracts.AccountInterface
}

func New(account contracts.AccountInterface) *Authorizer {
	return &Authorizer{account: account}
}
