package contact

import "github.com/dysodeng/wx/kernel/contracts"

// Contact 通讯录管理
type Contact struct {
	account contracts.AccountInterface
}

func New(account contracts.AccountInterface) *Contact {
	return &Contact{account: account}
}

// User 通讯录管理-成员管理
func (b *Contact) User() *User {
	return NewUser(b.account)
}

// Department 通讯录管理-部门管理
func (b *Contact) Department() *Department {
	return NewDepartment(b.account)
}
