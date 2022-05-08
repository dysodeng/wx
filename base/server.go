package base

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/dysodeng/wx/kernel/event"

	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/kernel/message"

	"github.com/dysodeng/wx/support/encryptor"
)

const (
	SuccessEmptyResponse = "success"
	EchoStr              = "echostr"
)

// Server 公众账号服务端
type Server struct {
	lock    sync.RWMutex
	account contracts.AccountInterface
	handler map[event.Guard]contracts.EventHandlerInterface
}

func NewServer(account contracts.AccountInterface) *Server {
	return &Server{
		account: account,
		handler: make(map[event.Guard]contracts.EventHandlerInterface),
	}
}

// Push 添加消息处理器
func (sg *Server) Push(handler contracts.EventHandlerInterface, guard event.Guard) {
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

	encrypt := encryptor.NewEncryptor(sg.account.AppId(), sg.account.Token(), sg.account.AesKey())
	if !encrypt.ValidSignature(timestamp, nonce, signature) {
		log.Println("signature is invalid")
		return
	}

	if e := request.FormValue(EchoStr); e != "" {
		_, _ = writer.Write([]byte(e))
		return
	}

	if request.Method == "POST" {

		var messageBody *message.Message

		if encryptType == "aes" {
			encryptRequestBody, err := encrypt.ParseEncryptBody(request)
			if err != nil {
				log.Printf("parse encrypt error: %+v", err)
				return
			}

			// Validate msg signature
			if !encrypt.ValidMsgSignature(timestamp, nonce, encryptRequestBody.Encrypt, msgSignature) {
				log.Println("msg_signature is invalid")
				return
			}

			// Decode base64
			cipherData, err := base64.StdEncoding.DecodeString(encryptRequestBody.Encrypt)
			if err != nil {
				log.Println("Decode base64 error:", err)
				return
			}

			// AES Decrypt
			plainData, err := encrypt.AesDecrypt(cipherData)
			if err != nil {
				fmt.Println("Aes decrypt error:", err)
				return
			}

			messageBody, _ = encrypt.ParseEncryptTextBody(plainData)
		} else {
			messageBody, _ = encrypt.ParseTextBody(request)
		}

		if messageBody.MsgType == "" && messageBody.InfoType != "" {
			messageBody.MsgType = messageBody.InfoType
		}

		var handler contracts.EventHandlerInterface
		var ok bool

		if messageBody.MsgType == "event" {
			if handler, ok = sg.handler[event.Event]; !ok {
				if handler, ok = sg.handler[event.Guard(strings.ToLower(messageBody.Event))]; !ok {
					handler, _ = sg.handler[event.All]
				}
			}
		} else {
			if handler, ok = sg.handler[event.Guard(messageBody.MsgType)]; !ok {
				handler, _ = sg.handler[event.All]
			}
		}

		if handler != nil {
			reply := handler.Handle(sg.account, messageBody)
			if reply != nil {
				replier := reply.Replier()
				xmlBody := replier.BuildXml(messageBody.ToUserName, messageBody.FromUserName)

				var replyBody []byte
				if encryptType == "aes" {
					replyBody, _ = encrypt.MakeEncryptBody(xmlBody, timestamp, nonce)
				} else {
					replyBody = xmlBody
				}

				writer.Header().Set("Content-Type", replier.ContentType())
				_, _ = writer.Write(replyBody)
				return
			}
		}
	}

	_, _ = writer.Write([]byte(SuccessEmptyResponse))
}
