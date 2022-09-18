package open_platform

import (
	"fmt"
	"time"

	"github.com/dysodeng/wx/kernel/message/reply"

	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/kernel/message"
)

// componentVerifyTicket component_verify_ticket 推送事件
type componentVerifyTicket struct{}

func (componentVerifyTicket) Handle(account contracts.AccountInterface, messageBody *message.Message) *reply.Reply {
	event := messageBody.EventMessage()
	verifyTicket := event.ComponentVerifyTicket()
	if verifyTicket.ComponentVerifyTicket != "" {
		cache, cacheKeyPrefix := account.Cache()
		cacheKey := cacheKeyPrefix + fmt.Sprintf(componentVerifyTicketCacheKey, verifyTicket.AppId)
		_ = cache.Put(cacheKey, verifyTicket.ComponentVerifyTicket, time.Second*42600)
	}
	return nil
}
