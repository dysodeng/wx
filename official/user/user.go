package user

import "github.com/dysodeng/wx/base"

// User 用户管理
type User struct {
	accessToken base.AccessTokenInterface
}

func NewUser(accessToken base.AccessTokenInterface) *User {
	return &User{accessToken: accessToken}
}
