package base

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/dysodeng/wx/base/encryptor"
	"github.com/dysodeng/wx/base/message"
)

const SuccessEmptyResponse = "success"

// Guard 服务类型
type Guard string

const (
	GuardAll     = "*"
	GuardEchoStr = "echostr"
)

// MessageReply 消息回复
type MessageReply struct {
	replier message.Replier
}

func NewMessageReply(replier message.Replier) *MessageReply {
	return &MessageReply{
		replier: replier,
	}
}

// GuardHandler 服务端消息处理器
type GuardHandler func(messageBody *message.Message) *MessageReply

// Server 公众账号服务端
type Server struct {
	lock    sync.RWMutex
	account AccountInterface
	request *http.Request
	writer  http.ResponseWriter
	handler map[Guard]GuardHandler
}

func NewServer(account AccountInterface, req *http.Request, writer http.ResponseWriter) *Server {
	return &Server{account: account, request: req, writer: writer, handler: make(map[Guard]GuardHandler)}
}

// Push 添加消息处理器
func (sg *Server) Push(handler GuardHandler, guard Guard) {
	sg.lock.Lock()
	defer sg.lock.Unlock()
	sg.handler[guard] = handler
}

func (sg *Server) Serve() {
	_ = sg.request.ParseForm()

	timestamp := strings.Join(sg.request.Form["timestamp"], "")
	nonce := strings.Join(sg.request.Form["nonce"], "")
	signature := strings.Join(sg.request.Form["signature"], "")
	encryptType := strings.Join(sg.request.Form["encrypt_type"], "")
	msgSignature := strings.Join(sg.request.Form["msg_signature"], "")

	encrypt := encryptor.NewEncryptor(sg.account.AccountAppId(), sg.account.AccountToken(), sg.account.AccountAesKey())
	if !encrypt.ValidSignature(timestamp, nonce, signature) {
		log.Println("Wechat Service: signature is invalid")
		return
	}

	if e := sg.request.FormValue("echostr"); e != "" {
		_, _ = sg.writer.Write([]byte(e))
		return
	}

	if sg.request.Method == "POST" {
		if encryptType == "aes" {

			encryptRequestBody, err := encrypt.ParseEncryptBody(sg.request)

			if err == nil {
				// Validate msg signature
				if !encrypt.ValidMsgSignature(timestamp, nonce, encryptRequestBody.Encrypt, msgSignature) {
					log.Println("Wechat Service: msg_signature is invalid")
					return
				}

				// Decode base64
				cipherData, err := base64.StdEncoding.DecodeString(encryptRequestBody.Encrypt)
				if err != nil {
					log.Println("Wechat Service: Decode base64 error:", err)
					return
				}

				// AES Decrypt
				plainData, err := encrypt.AesDecrypt(cipherData)
				if err != nil {
					fmt.Println(err)
					return
				}

				messageBody, _ := encrypt.ParseEncryptTextBody(plainData)
				log.Println(messageBody)

				var handler GuardHandler
				var ok bool

				sg.lock.RLock()
				if handler, ok = sg.handler[Guard(messageBody.MsgType)]; !ok {
					if handler, ok = sg.handler[GuardAll]; ok {
					}
				}
				sg.lock.RUnlock()

				if handler != nil {
					reply := handler(messageBody)
					if reply != nil {

						xmlBody := reply.replier.BuildXml(messageBody.ToUserName, messageBody.FromUserName)
						replyBody, _ := encrypt.MakeEncryptBody(xmlBody, timestamp, nonce)

						sg.writer.Header().Set("Content-Type", reply.replier.ContentType())
						_, _ = sg.writer.Write(replyBody)

						return
					}
				}
			}
		}
	}

	_, _ = sg.writer.Write([]byte(SuccessEmptyResponse))
}
