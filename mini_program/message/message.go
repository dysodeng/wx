package message

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// Message 小程序订阅消息
type Message struct {
	account contracts.AccountInterface
}

// Body 消息体
type Body struct {
	ToUser           string                `json:"touser"`
	TemplateId       string                `json:"template_id"`
	Path             string                `json:"path"`
	Data             map[string]*DataValue `json:"data,omitempty"`
	MiniProgramState string                `json:"miniprogram_state"`
	Lang             string                `json:"lang"`
}

type DataValue struct {
	Value string `json:"value"`
}

func New(account contracts.AccountInterface) *Message {
	return &Message{account: account}
}

// Send 发送订阅消息
func (m *Message) Send(message Body) error {
	if message.ToUser == "" {
		return kernelError.New(0, errors.New("attribute touser can not be empty!"))
	}
	if message.TemplateId == "" {
		return kernelError.New(0, errors.New("attribute template_id can not be empty!"))
	}
	accessToken, _ := m.account.AccessToken()
	apiUrl := fmt.Sprintf(
		"cgi-bin/message/subscribe/send?access_token=%s",
		accessToken.AccessToken,
	)

	if message.MiniProgramState == "" {
		message.MiniProgramState = "formal"
	}
	if message.Lang == "" {
		message.Lang = "zh_CN"
	}

	res, err := http.PostJSON(apiUrl, message)
	if err != nil {
		return kernelError.New(0, err)
	}

	var result kernelError.ApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return kernelError.New(0, err)
	}

	return kernelError.NewWithApiError(result)
}
