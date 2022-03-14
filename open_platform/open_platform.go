package open_platform

import (
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
func (open *OpenPlatform) Server() *base.Server {
	return base.NewServer(open)
}
