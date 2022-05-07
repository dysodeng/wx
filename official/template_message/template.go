package template_message

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"

	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

const defaultColor = "#000000"

// TemplateMessage 模板消息
type TemplateMessage struct {
	accessToken contracts.AccessTokenInterface
}

func NewTemplateMessage(accessToken contracts.AccessTokenInterface) *TemplateMessage {
	return &TemplateMessage{accessToken: accessToken}
}

// Industry 所属行业
type Industry struct {
	PrimaryIndustry   IndustryItem `json:"primary_industry"`
	SecondaryIndustry IndustryItem `json:"secondary_industry"`
}
type IndustryItem struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

type TemplateList struct {
	TemplateList []Template `json:"template_list"`
}

// Template 模板信息
type Template struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

// Message 消息体
type Message struct {
	ToUser      string                       `json:"touser"`
	TemplateId  string                       `json:"template_id"`
	TopColor    string                       `json:"topcolor"`
	Url         string                       `json:"url"`
	MiniProgram *MiniProgram                 `json:"miniprogram,omitempty"`
	Data        map[string]*MessageDataValue `json:"data,omitempty"`
}

// MiniProgram 跳转小程序
type MiniProgram struct {
	AppID    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}

// MessageDataValue 消息数据值
type MessageDataValue struct {
	Value string `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}

// SetIndustry 设置所属行业
func (tm *TemplateMessage) SetIndustry(industryOne, industryTwo string) error {
	accessToken, _ := tm.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf(
		"cgi-bin/template/api_set_industry?access_token=%s",
		accessToken.AccessToken,
	)

	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"industry_id1": industryOne,
		"industry_id2": industryTwo,
	})
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

// GetIndustry 获取设置的行业信息
func (tm *TemplateMessage) GetIndustry() (*Industry, error) {
	accessToken, _ := tm.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/template/get_industry?access_token=%s", accessToken.AccessToken)

	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, baseError.New(0, err)
	}

	// 返回信息
	type industryResult struct {
		baseError.WxApiError
		Industry
	}
	var result industryResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return &result.Industry, nil
}

// AddTemplate 添加模板
func (tm *TemplateMessage) AddTemplate(templateIdShort string) (string, error) {
	accessToken, _ := tm.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/template/api_add_template?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"template_id_short": templateIdShort})
	if err != nil {
		return "", baseError.New(0, err)
	}

	// 返回信息
	type tempResult struct {
		baseError.WxApiError
		TemplateId string `json:"template_id"`
	}

	var result tempResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return "", baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.TemplateId, nil
}

// GetPrivateTemplates 获取所有模板列表
func (tm *TemplateMessage) GetPrivateTemplates() (*TemplateList, error) {
	accessToken, _ := tm.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/template/get_all_private_template?access_token=%s", accessToken.AccessToken)

	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, baseError.New(0, err)
	}

	// 返回信息
	type templateListResult struct {
		baseError.WxApiError
		TemplateList
	}
	var result templateListResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return &result.TemplateList, nil
}

// DeletePrivateTemplate 删除模板
func (tm *TemplateMessage) DeletePrivateTemplate(templateId string) error {
	accessToken, _ := tm.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/template/del_private_template?access_token=%s", accessToken.AccessToken)

	res, err := http.PostJson(apiUrl, map[string]interface{}{"template_id": templateId})
	if err != nil {
		return baseError.New(0, err)
	}

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

// Send 模板消息发送
func (tm *TemplateMessage) Send(message Message) (int, error) {
	if message.ToUser == "" {
		return 0, baseError.New(0, errors.New("attribute touser can not be empty!"))
	}
	if message.TemplateId == "" {
		return 0, baseError.New(0, errors.New("attribute template_id can not be empty!"))
	}
	if message.TopColor == "" {
		message.TopColor = defaultColor
	}
	if message.Data != nil {
		for key, _ := range message.Data {
			if message.Data[key].Color == "" {
				message.Data[key].Color = defaultColor
			}
		}
	}

	accessToken, _ := tm.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf("cgi-bin/message/template/send?access_token=%s", accessToken.AccessToken)

	body, _ := json.Marshal(message)
	var messageBody map[string]interface{}
	_ = json.Unmarshal(body, &messageBody)

	res, err := http.PostJson(apiUrl, messageBody)
	if err != nil {
		return 0, baseError.New(0, err)
	}

	// 返回信息
	type sendResult struct {
		baseError.WxApiError
		MsgId int `json:"msgid"`
	}

	var result sendResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return 0, baseError.New(0, err)
	}
	if result.ErrCode != 0 {
		return 0, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.MsgId, nil
}
