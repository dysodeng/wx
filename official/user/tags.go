package user

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"

	"github.com/dysodeng/wx/support/http"
)

// TagItem 用户标签
type TagItem struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type TagUser struct {
	Count int `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

// Tag 用户标签管理
type Tag struct {
	accessToken contracts.AccessTokenInterface
}

func NewUserTag(accessToken contracts.AccessTokenInterface) *Tag {
	return &Tag{accessToken: accessToken}
}

// Create 创建标签
func (tag *Tag) Create(name string) (TagItem, error) {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/tags/create?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"tag": map[string]string{"name": name}})
	if err != nil {
		return TagItem{}, baseError.New(0, err)
	}

	// 返回信息
	type tagResult struct {
		baseError.WxApiError
		Tag TagItem `json:"tag"`
	}
	var result tagResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return TagItem{}, baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return TagItem{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.Tag, nil
}

// List 获取公众号已创建的标签
func (tag *Tag) List() ([]TagItem, error) {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/tags/get?access_token=%s", accessToken.AccessToken)

	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, baseError.New(0, err)
	}

	// 返回信息
	type tagResult struct {
		baseError.WxApiError
		Tags []TagItem `json:"tags"`
	}
	var result tagResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.Tags, nil
}

// Update 编辑标签
func (tag *Tag) Update(tagId int, name string) error {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/tags/update?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"tag": map[string]interface{}{"id": tagId, "name": name}})
	if err != nil {
		return baseError.New(0, err)
	}

	// 返回信息
	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// Delete 删除标签
func (tag *Tag) Delete(tagId int) error {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/tags/delete?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"tag": map[string]interface{}{"id": tagId}})
	if err != nil {
		return baseError.New(0, err)
	}

	// 返回信息
	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// UsersOfTag 获取标签下粉丝列表
func (tag *Tag) UsersOfTag(tagId int, nextOpenid string) (TagUser, error) {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/user/tag/get?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"tagid": tagId, "next_openid": nextOpenid})
	if err != nil {
		return TagUser{}, baseError.New(0, err)
	}

	// 返回信息
	type tagResult struct {
		baseError.WxApiError
		TagUser
	}
	var result tagResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return TagUser{}, baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return TagUser{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return TagUser{Count: result.Count, Data: result.Data, NextOpenid: result.NextOpenid}, nil
}

// TagUsers 批量为用户打标签
func (tag *Tag) TagUsers(openidList []string, tagId int) error {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/tags/members/batchtagging?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"openid_list": openidList, "tagid": tagId})
	if err != nil {
		return baseError.New(0, err)
	}

	// 返回信息
	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// UntagUsers 批量为用户取消标签
func (tag *Tag) UntagUsers(openidList []string, tagId int) error {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/tags/members/batchuntagging?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"openid_list": openidList, "tagid": tagId})
	if err != nil {
		return baseError.New(0, err)
	}

	// 返回信息
	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// UserTags 获取用户身上的标签列表
func (tag *Tag) UserTags(openid string) ([]int, error) {
	accessToken, _ := tag.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/tags/getidlist?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"openid": openid})
	if err != nil {
		return nil, baseError.New(0, err)
	}

	// 返回信息
	type tagResult struct {
		baseError.WxApiError
		TagIdList []int `json:"tagid_list"`
	}
	var result tagResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.TagIdList, nil
}
