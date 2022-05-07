package open_platform

import (
	"fmt"
	"log"
	"time"

	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/kernel/message"
)

const componentVerifyTicketCacheKey = "component_verify_ticket.%s"

// ComponentVerifyTicket verify_ticket推送事件
type ComponentVerifyTicket struct{}

func (ComponentVerifyTicket) Handle(account contracts.AccountInterface, messageBody *message.Message) *message.Reply {
	log.Printf("appId:%s, component_verify_ticket:%s", messageBody.AppId, messageBody.ComponentVerifyTicket)
	log.Println("current server appId:%", account.AccountAppId())
	if messageBody.ComponentVerifyTicket != "" {
		cache, cacheKeyPrefix := account.Cache()
		cacheKey := cacheKeyPrefix + fmt.Sprintf(componentVerifyTicketCacheKey, messageBody.AppId)
		_ = cache.Put(cacheKey, messageBody.ComponentVerifyTicket, time.Second*42600)
	}
	return nil
}
