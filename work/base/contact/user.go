package contact

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/work/http"
)

// User 通讯录成员管理
type User struct {
	account contracts.AccountInterface
}

func NewUser(account contracts.AccountInterface) *User {
	return &User{account: account}
}

// Create 创建成员
func (u *User) Create(user CreateUserRequest) error {
	accessToken, err := u.account.AccessToken()
	if err != nil {
		return kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/user/create?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, user)
	if err != nil {
		return kernelError.New(0, err)
	}

	var result kernelError.ApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return kernelError.NewWithApiError(result)
	}

	return nil
}

// Get 读取成员
func (u *User) Get(userid string) (*UserInfo, error) {
	accessToken, err := u.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/user/get?access_token=%s&userid=%s", accessToken.AccessToken, userid)
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result userInfoResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.UserInfo, nil
}

// Update 更新成员
func (u *User) Update(user UpdateUserRequest) error {
	accessToken, err := u.account.AccessToken()
	if err != nil {
		return kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/user/update?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, user)
	if err != nil {
		return kernelError.New(0, err)
	}

	var result kernelError.ApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return kernelError.NewWithApiError(result)
	}

	return nil
}

// Delete 删除成员
func (u *User) Delete(userid string) error {
	accessToken, err := u.account.AccessToken()
	if err != nil {
		return kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/user/delete?access_token=%s&userid=%s", accessToken.AccessToken, userid)
	res, err := http.Get(apiUrl)
	if err != nil {
		return kernelError.New(0, err)
	}

	var result kernelError.ApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return kernelError.NewWithApiError(result)
	}

	return nil
}

// BatchDelete 批量删除成员
func (u *User) BatchDelete(useridList []string) error {
	accessToken, err := u.account.AccessToken()
	if err != nil {
		return kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/user/batchdelete?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, map[string]interface{}{"useridlist": useridList})
	if err != nil {
		return kernelError.New(0, err)
	}

	var result kernelError.ApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return kernelError.NewWithApiError(result)
	}

	return nil
}

// ListId 获取成员ID列表
func (u *User) ListId(cursor string, limit int) (*UserIdList, error) {
	accessToken, err := u.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/user/list_id?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
	})
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result userIdListResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.UserIdList, nil
}
