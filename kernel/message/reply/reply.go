package reply

// Replier 消息回复接口
type Replier interface {
	// BuildXml 构建消息XML
	BuildXml(fromUserName, toUserName string) []byte
	// ContentType 消息类型
	ContentType() string
}

// Reply 消息回复
type Reply struct {
	replier Replier
}

func NewReply(replier Replier) *Reply {
	return &Reply{
		replier: replier,
	}
}

func (reply *Reply) Replier() Replier {
	return reply.replier
}
