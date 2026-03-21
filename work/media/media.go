package media

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/work/http"
)

// Media 临时素材管理
type Media struct {
	account contracts.AccountInterface
}

func NewMedia(account contracts.AccountInterface) *Media {
	return &Media{account: account}
}

// Upload 上传临时素材
// mediaType: image-图片, voice-语音, video-视频, file-文件
func (m *Media) Upload(mediaType string, filename string, fileData []byte) (*UploadResult, error) {
	accessToken, err := m.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/media/upload?access_token=%s&type=%s", accessToken.AccessToken, mediaType)
	res, err := http.Upload(apiUrl, filename, fileData)
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result uploadResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.UploadResult, nil
}
