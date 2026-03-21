package media

import kernelError "github.com/dysodeng/wx/kernel/error"

// UploadResult 上传临时素材结果
type UploadResult struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// uploadResponse 上传临时素材响应
type uploadResponse struct {
	kernelError.ApiError
	UploadResult
}
