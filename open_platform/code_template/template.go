package code_template

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// CodeTemplate 小程序代码模板
type CodeTemplate struct {
	account contracts.AccountInterface
}

func NewCodeTemplate(account contracts.AccountInterface) *CodeTemplate {
	return &CodeTemplate{account: account}
}

// GetDraftList 获取代码草稿列表
func (template *CodeTemplate) GetDraftList() ([]Draft, error) {
	accountToken, err := template.account.AccessToken()
	if err != nil {
		return nil, err
	}

	apiUrl := fmt.Sprintf("wxa/gettemplatedraftlist?access_token=%s", accountToken.AccessToken)
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	var result struct {
		baseError.WxApiError
		DraftList []Draft `json:"draft_list"`
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	if err == nil && result.ErrCode != 0 {
		return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.DraftList, nil
}

// AddDraftToTemplate 将草稿添加到代码模板库
// draftId 草稿ID
// templateType 模板类型 0-普通模板 1-标准模板
func (template *CodeTemplate) AddDraftToTemplate(draftId int64, templateType uint8) error {
	accountToken, err := template.account.AccessToken()
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("wxa/addtotemplate?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"draft_id":      draftId,
		"template_type": templateType,
	})
	if err != nil {
		return baseError.New(0, err)
	}

	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return err
	}
	if err == nil && result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// GetTemplateList 获取代码模板列表
// templateType 模板类型 -1:所有模板 0-普通模板 1-标准模板
func (template *CodeTemplate) GetTemplateList(templateType int8) ([]Template, error) {
	accountToken, err := template.account.AccessToken()
	if err != nil {
		return nil, err
	}

	var apiUrl string
	if templateType == -1 {
		apiUrl = fmt.Sprintf("wxa/gettemplatedraftlist?access_token=%s", accountToken.AccessToken)
	} else {
		apiUrl = fmt.Sprintf("wxa/gettemplatedraftlist?access_token=%stemplate_type=%d", accountToken.AccessToken, templateType)
	}

	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	var result struct {
		baseError.WxApiError
		TemplateList []Template `json:"template_list"`
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	if err == nil && result.ErrCode != 0 {
		return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.TemplateList, nil
}

func (template *CodeTemplate) DeleteTemplate(templateId int64) error {
	accountToken, err := template.account.AccessToken()
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("wxa/addtotemplate?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"template_id": templateId,
	})
	if err != nil {
		return baseError.New(0, err)
	}

	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return err
	}
	if err == nil && result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// Draft 模板草稿
type Draft struct {
	CreateTime  int64  `json:"create_time"`
	DraftId     int64  `json:"draft_id"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
}

// Template 模板信息
type Template struct {
	TemplateId             int64                    `json:"template_id"`
	TemplateType           uint8                    `json:"template_type"`
	CreateTime             int64                    `json:"create_time"`
	UserVersion            string                   `json:"user_version"`
	UserDesc               string                   `json:"user_desc"`
	SourceMiniProgramAppid string                   `json:"source_miniprogram_appid"`
	SourceMiniProgram      string                   `json:"source_miniprogram"`
	Developer              string                   `json:"developer"`
	AuditScene             int8                     `json:"audit_scene"`
	AuditStatus            int8                     `json:"audit_status"`
	Reason                 string                   `json:"reason"`
	CategoryList           []map[string]interface{} `json:"category_list"`
}
