package user

import (
	"github.com/dysodeng/wx/kernel/contracts"
)

// User 用户管理
type User struct {
	accessToken contracts.AccessTokenInterface
}

func NewUser(accessToken contracts.AccessTokenInterface) *User {
	return &User{accessToken: accessToken}
}
