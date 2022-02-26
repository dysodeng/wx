package open_platform

type OpenPlatform struct {
}

func NewOpenPlatform() *OpenPlatform {
	return &OpenPlatform{}
}

func (op OpenPlatform) AccessToken(refresh bool) string {
	return ""
}
