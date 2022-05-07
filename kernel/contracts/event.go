package contracts

import (
	"github.com/dysodeng/wx/kernel/message"
)

// EventHandlerInterface 事件处理器接口
type EventHandlerInterface interface {
	Handle(account AccountInterface, messageBody *message.Message) *message.Reply
}
