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

// GuardHandler 服务端消息处理器
type GuardHandler func(messageBody *message.Message) *message.Reply

// Server 公众账号服务端
type Server struct {
	lock    sync.RWMutex
	account AccountInterface
	handler map[Guard]GuardHandler
}

func NewServer(account AccountInterface) *Server {
	return &Server{
		account: account,
		handler: make(map[Guard]GuardHandler),
	}
}

// Push 添加消息处理器
func (sg *Server) Push(handler GuardHandler, guard Guard) {
	sg.lock.Lock()
	defer sg.lock.Unlock()

	sg.handler[guard] = handler
}

// Serve Handle and return response.
func (sg *Server) Serve(request *http.Request, writer http.ResponseWriter) {
	_ = request.ParseForm()

	timestamp := strings.Join(request.Form["timestamp"], "")
	nonce := strings.Join(request.Form["nonce"], "")
	signature := strings.Join(request.Form["signature"], "")
	encryptType := strings.Join(request.Form["encrypt_type"], "")
	msgSignature := strings.Join(request.Form["msg_signature"], "")

	encrypt := encryptor.NewEncryptor(sg.account.AccountAppId(), sg.account.AccountToken(), sg.account.AccountAesKey())
	if !encrypt.ValidSignature(timestamp, nonce, signature) {
		log.Println("Wechat Service: signature is invalid")
		return
	}

	if e := request.FormValue("echostr"); e != "" {
		_, _ = writer.Write([]byte(e))
		return
	}

	if request.Method == "POST" {
		if encryptType == "aes" {

			encryptRequestBody, err := encrypt.ParseEncryptBody(request)

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

				handlerName := ""
				if messageBody.MsgType == "event" {
					handlerName = strings.ToLower(messageBody.Event)
				} else {
					handlerName = messageBody.MsgType
				}

				sg.lock.RLock()
				if handler, ok = sg.handler[Guard(handlerName)]; !ok {
					handler, _ = sg.handler[GuardAll]
				}
				sg.lock.RUnlock()

				if handler != nil {
					reply := handler(messageBody)
					if reply != nil {

						replier := reply.Replier()

						xmlBody := replier.BuildXml(messageBody.ToUserName, messageBody.FromUserName)
						replyBody, _ := encrypt.MakeEncryptBody(xmlBody, timestamp, nonce)

						writer.Header().Set("Content-Type", replier.ContentType())
						_, _ = writer.Write(replyBody)

						return
					}
				}
			}
		}
	}

	_, _ = writer.Write([]byte(SuccessEmptyResponse))
}
