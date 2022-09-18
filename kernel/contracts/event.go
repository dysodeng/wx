package contracts

import (
	"github.com/dysodeng/wx/kernel/message"
	"github.com/dysodeng/wx/kernel/message/reply"
)

// EventHandlerInterface 事件处理器接口
type EventHandlerInterface interface {
	Handle(account AccountInterface, messageBody *message.Message) *reply.Reply
}
