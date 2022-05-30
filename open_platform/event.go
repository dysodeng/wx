package open_platform

import (
	"fmt"
	"time"

	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/kernel/message"
)

// ComponentVerifyTicket verify_ticket推送事件
type ComponentVerifyTicket struct{}

func (ComponentVerifyTicket) Handle(account contracts.AccountInterface, messageBody *message.Message) *message.Reply {
	if messageBody.ComponentVerifyTicket != "" {
		cache, cacheKeyPrefix := account.Cache()
		cacheKey := cacheKeyPrefix + fmt.Sprintf(componentVerifyTicketCacheKey, messageBody.AppId)
		_ = cache.Put(cacheKey, messageBody.ComponentVerifyTicket, time.Second*42600)
	}
	return nil
}
