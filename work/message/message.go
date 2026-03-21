package message

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/work/http"
)

// Message 应用消息
type Message struct {
	account contracts.AccountInterface
}

func NewMessage(account contracts.AccountInterface) *Message {
	return &Message{account: account}
}

// Chat 群聊会话
func (m *Message) Chat() *Chat {
	return newChat(m)
}

// Send 发送应用消息
func (m *Message) Send(req SendRequest) (*SendResult, error) {
	accessToken, err := m.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/message/send?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, req)
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result sendResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.SendResult, nil
}

// UpdateTemplateCard 更新模版卡片消息
func (m *Message) UpdateTemplateCard(req UpdateTemplateCardRequest) (*UpdateTemplateCardResult, error) {
	accessToken, err := m.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/message/update_template_card?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, req)
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result updateTemplateCardResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.UpdateTemplateCardResult, nil
}

// Recall 撤回应用消息
func (m *Message) Recall(msgId string) error {
	accessToken, err := m.account.AccessToken()
	if err != nil {
		return kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/message/recall?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, RecallRequest{MsgId: msgId})
	if err != nil {
		return kernelError.New(0, err)
	}

	var result recallResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return kernelError.NewWithApiError(result.ApiError)
	}

	return nil
}
