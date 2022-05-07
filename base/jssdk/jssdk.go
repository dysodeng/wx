package jssdk

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dysodeng/wx/base"
	baseError "github.com/dysodeng/wx/base/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/dysodeng/wx/support/str"
)

const cacheKeyTemplate = "jssdk.ticket.%s.%s"

// Jssdk 微信JSSDK
type Jssdk struct {
	accessToken base.AccountInterface
	url         string
}

type signatureConfig struct {
	nonce     string
	timestamp int64
	url       string
	signature string
}

type Ticket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func NewJssdk(accessToken base.AccountInterface) *Jssdk {
	return &Jssdk{accessToken: accessToken}
}

func (js *Jssdk) SetUrl(url string) *Jssdk {
	js.url = url
	return js
}

// BuildConfig 构建jssdk配置信息
func (js *Jssdk) BuildConfig(jsApiList []string, debug, beta bool) map[string]interface{} {
	config := js.configSignature()
	return map[string]interface{}{
		"debug":     debug,
		"beta":      beta,
		"jsApiList": jsApiList,
		"appId":     js.accessToken.AccountAppId(),
		"nonceStr":  config.nonce,
		"timestamp": config.timestamp,
		"url":       config.url,
		"signature": config.signature,
	}
}

// getTicket 获取ticket ticketType:ticket类型 jsapi与wx_card
func (js *Jssdk) getTicket(ticketType string) Ticket {
	cache, cacheKeyPrefix := js.accessToken.Cache()
	cacheKey := cacheKeyPrefix + fmt.Sprintf(cacheKeyTemplate, ticketType, js.accessToken.AccountAppId())

	if cache.IsExist(cacheKey) {
		ticketString, err := cache.Get(cacheKey)
		if err == nil {
			if t, ok := ticketString.(string); ok {
				var ticketBody Ticket
				err = json.Unmarshal([]byte(t), &ticketBody)
				if err == nil {
					return ticketBody
				}
			}

		}
	}

	// 获取新的ticket
	return js.refreshTicket(ticketType)
}

// refreshTicket 刷新ticket
func (js *Jssdk) refreshTicket(ticketType string) Ticket {
	accessToken, _ := js.accessToken.AccessToken(false)
	apiUrl := fmt.Sprintf(
		"cgi-bin/ticket/getticket?access_token=%s&type=%s",
		accessToken.AccessToken,
		ticketType,
	)
	res, err := http.Get(apiUrl)
	if err != nil {
		return Ticket{}
	}

	// 返回信息
	type ticketResult struct {
		baseError.WxApiError
		Ticket
	}
	var result ticketResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return Ticket{}
	}

	if result.ErrCode != 0 {
		return Ticket{}
	}

	ticketByte, _ := json.Marshal(Ticket{
		Ticket:    result.Ticket.Ticket,
		ExpiresIn: result.Ticket.ExpiresIn,
	})

	// 缓存
	cache, cacheKeyPrefix := js.accessToken.Cache()
	cacheKey := cacheKeyPrefix + fmt.Sprintf(cacheKeyTemplate, ticketType, js.accessToken.AccountAppId())

	_ = cache.Put(
		cacheKey,
		string(ticketByte),
		time.Second*time.Duration(result.Ticket.ExpiresIn-600), // 提前过期
	)

	return result.Ticket
}

// configSignature jssdk 配置签名
func (js *Jssdk) configSignature() signatureConfig {
	nonce := str.RandString(10)
	timestamp := time.Now().Unix()
	ticket := js.getTicket("jsapi")
	return signatureConfig{
		nonce:     nonce,
		timestamp: timestamp,
		url:       js.url,
		signature: js.getTicketSignature(ticket.Ticket, nonce, timestamp, js.url),
	}
}

func (js *Jssdk) getTicketSignature(ticket, nonce string, timestamp int64, url string) string {
	sha := sha1.New()
	sha.Write([]byte(fmt.Sprintf(
		"jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s",
		ticket,
		nonce,
		timestamp,
		url,
	)))
	return hex.EncodeToString(sha.Sum(nil))
}
