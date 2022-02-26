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
	"github.com/dysodeng/wx/official"
	"log"
)

var (
	appId     string = ""
	appSecret string = ""
	token     string = ""
	aesKey    string = ""
)

func main() {
	officialSdk, err := official.NewOfficial(official.WithOfficial(appId, appSecret, token, aesKey))
	if err != nil {
		log.Fatal(err)
	}

	userTag := officialSdk.UserTag()
	log.Println(userTag.List())
}
```
