package contracts

import (
	"context"

	"github.com/dysodeng/wx/kernel/message"
	"github.com/dysodeng/wx/kernel/message/reply"
)

// EventHandler 事件处理函数
type EventHandler func(ctx context.Context, account AccountInterface, msg *message.Message) (*reply.Reply, error)

// Middleware 中间件
type Middleware func(next EventHandler) EventHandler
