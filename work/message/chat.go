package message

import (
	"encoding/json"
	"fmt"

	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/work/http"
)

// Chat 群聊会话
type Chat struct {
	message *Message
}

func newChat(message *Message) *Chat {
	return &Chat{message: message}
}

// Create 创建群聊会话
func (c *Chat) Create(req CreateChatRequest) (string, error) {
	accessToken, err := c.message.account.AccessToken()
	if err != nil {
		return "", kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/appchat/create?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, req)
	if err != nil {
		return "", kernelError.New(0, err)
	}

	var result createChatResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return "", kernelError.NewWithApiError(result.ApiError)
	}

	return result.ChatId, nil
}

// Update 修改群聊会话
func (c *Chat) Update(req UpdateChatRequest) error {
	accessToken, err := c.message.account.AccessToken()
	if err != nil {
		return kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/appchat/update?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, req)
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

// Get 获取群聊会话
func (c *Chat) Get(chatId string) (*ChatInfo, error) {
	accessToken, err := c.message.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/appchat/get?access_token=%s&chatid=%s", accessToken.AccessToken, chatId)
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result getChatResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.ChatInfo, nil
}

// Send 应用推送消息到群聊会话
func (c *Chat) Send(req ChatSendRequest) error {
	accessToken, err := c.message.account.AccessToken()
	if err != nil {
		return kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/appchat/send?access_token=%s", accessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, req)
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
