package open_platform

import (
	"net/http"

	"github.com/dysodeng/wx/base"
)

// OpenPlatform 开放平台
type OpenPlatform struct {
	config *config
	option *option
}

func NewOpenPlatform() *OpenPlatform {
	return &OpenPlatform{}
}

// Server 服务端
func (open *OpenPlatform) Server(req *http.Request, writer http.ResponseWriter) *base.Server {
	return base.NewServer(open, req, writer)
}
