package user

import (
	"github.com/dysodeng/wx/kernel/contracts"
)

// User 用户管理
type User struct {
	account contracts.AccountInterface
}

func NewUser(account contracts.AccountInterface) *User {
	return &User{account: account}
}
