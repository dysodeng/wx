package mini_program

import (
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
)

func (w *MiniProgram) AccessToken() (contracts.AccessToken, error) {
	return contracts.AccessToken{}, nil
}

func (w *MiniProgram) AccessTokenCacheKey() string {
	return fmt.Sprintf("%s%s.%s", w.option.cacheKeyPrefix, "access_token", w.config.corpId)
}
