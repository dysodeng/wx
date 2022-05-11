package content

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// Security 内容安全
type Security struct {
	account contracts.AccountInterface
}

func NewContentSecurity(account contracts.AccountInterface) *Security {
	return &Security{account: account}
}

// CheckText 文本内容审核
func (c *Security) CheckText(openid, content string, scene Scene) (TextResult, error) {
	accountToken, err := c.account.AccessToken()
	if err != nil {
		return TextResult{}, err
	}

	apiUrl := fmt.Sprintf("wxa/msg_sec_check?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"version": 2,
		"openid":  openid,
		"scene":   scene,
		"content": content,
	})
	if err != nil {
		return TextResult{}, baseError.New(0, err)
	}

	var result struct {
		baseError.WxApiError
		TextResult
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return TextResult{}, err
	}
	if err == nil && result.ErrCode != 0 {
		return TextResult{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.TextResult, nil
}

// AsyncCheckMedia 异步校验图片/音频是否合规
func (c *Security) AsyncCheckMedia(openid, mediaUrl string, mediaType MediaType, scene Scene) (string, error) {
	accountToken, err := c.account.AccessToken()
	if err != nil {
		return "", err
	}

	apiUrl := fmt.Sprintf("wxa/media_check_async?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"version":    2,
		"openid":     openid,
		"scene":      scene,
		"media_url":  mediaUrl,
		"media_type": mediaType,
	})
	if err != nil {
		return "", baseError.New(0, err)
	}
	var result struct {
		baseError.WxApiError
		TraceId string `json:"trace_id"`
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}
	if err == nil && result.ErrCode != 0 {
		return "", baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.TraceId, nil
}

// Scene 审核场景
type Scene int

const (
	Material  Scene = 1 // 资料
	Comment   Scene = 2 // 评论
	Forum     Scene = 3 // 论坛
	SocialLog Scene = 4 // 社交日志
)

// MediaType 媒体类型
type MediaType uint8

const (
	Audio MediaType = 1
	Image MediaType = 2
)

type Result struct {
	Suggest string `json:"suggest"`
	Label   int    `json:"label"`
}

type Detail struct {
	Strategy string `json:"strategy"`
	ErrCode  int    `json:"errcode"`
	Suggest  string `json:"suggest"`
	Label    int    `json:"label"`
	Level    int    `json:"level"`
	Prob     int    `json:"prob"`
	Keyword  string `json:"keyword"`
}

// TextResult 文字审查结果
type TextResult struct {
	TraceId string   `json:"trace_id"`
	Result  Result   `json:"result"`
	Detail  []Detail `json:"detail"`
}

// TextSceneError 获取违规内容错误信息
func TextSceneError(scene int) string {
	var err string
	switch scene {
	case 10001:
		err = "广告内容"
	case 20001:
		err = "时政内容"
	case 20002:
		err = "色情内容"
	case 20003:
		err = "辱骂内容"
	case 20006:
		err = "违法犯罪内容"
	case 20008:
		err = "欺诈内容"
	case 20012:
		err = "低俗内容"
	case 20013:
		err = "版权内容"
	case 21000:
		err = "其它"
	}
	return err
}
