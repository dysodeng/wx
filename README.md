wx
==========
微信公众账号SDK

- 包含公众号、小程序与开放平台

Installation
------------
```sh
go get github.com/dysodeng/wx
```

Usage
-----
微信公众号调用接口

```golang
package main

import (
	"github.com/dysodeng/wx/kernel/event"
	"github.com/dysodeng/wx/kernel/message"
	"github.com/dysodeng/wx/kernel/message/reply"
	"github.com/dysodeng/wx/official"
	"log"
	"net/http"
)

var (
	appId     string = ""
	appSecret string = ""
	token     string = ""
	aesKey    string = ""
)

func main() {
	officialSdk := official.New(appId, appSecret, token, aesKey)

	// 公众号接口调用
	user := officialSdk.User()
	log.Println(user.Info("openid"))

	// 服务器端
	wxServer := officialSdk.Server()
	wxServer.Register(func(messageBody *message.Message) *reply.Reply {
		log.Println("这里是用户自定义的消息处理器")
		log.Println(messageBody)
		return reply.NewReply(reply.NewText("你好呀"))
	}, event.All)

	h := http.DefaultServeMux
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wxServer.Serve(r, w)
	})
	server := http.Server{
		Addr:    ":80",
		Handler: h,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("%+v\n", err)
	}
}
```
