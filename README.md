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
非开放平台下微信公众号调用sdk
```golang
package main

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/base/message"
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
	officialSdk, err := official.NewOfficial(official.WithAccount(appId, appSecret, token, aesKey))
	if err != nil {
		log.Fatal(err)
	}

	// 公众号接口调用
	userTag := officialSdk.UserTag()
	log.Println(userTag.List())
	
	// 服务端
	appServer := officialSdk.Server()
	appServer.Push(func(messageBody *message.Message) *message.Reply {
		log.Println("这里是用户自定义的消息处理器")
		log.Println(messageBody)
		return message.NewMessageReply(message.NewText("你好呀"))
	}, base.GuardAll)
	
	h := http.DefaultServeMux
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		appServer.Serve(r, w)
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
