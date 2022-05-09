package jssdk

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dysodeng/wx/support"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"

	"github.com/dysodeng/wx/support/http"
)

const cacheKeyTemplate = "jssdk.ticket.%s.%s"

// Jssdk 微信JSSDK
type Jssdk struct {
	account contracts.AccountInterface
	url     string
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

func NewJssdk(account contracts.AccountInterface) *Jssdk {
	return &Jssdk{account: account}
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
		"appId":     js.account.AppId(),
		"nonceStr":  config.nonce,
		"timestamp": config.timestamp,
		"url":       config.url,
		"signature": config.signature,
	}
}

// getTicket 获取ticket ticketType:ticket类型 jsapi与wx_card
func (js *Jssdk) getTicket(ticketType string) Ticket {
	cache, cacheKeyPrefix := js.account.Cache()
	cacheKey := cacheKeyPrefix + js.getTicketCacheKey(ticketType)

	if cache.IsExist(cacheKey) {
		ticketString, err := cache.Get(cacheKey)
		if err == nil {
			var ticketBody Ticket
			err = json.Unmarshal([]byte(ticketString), &ticketBody)
			if err == nil {
				return ticketBody
			}
		}
	}

	// 获取新的ticket
	return js.refreshTicket(ticketType)
}

// refreshTicket 刷新ticket
func (js *Jssdk) refreshTicket(ticketType string) Ticket {
	var accessToken string
	if js.account.IsOpenPlatform() {
		accessToken = js.account.ComponentAccessToken()
	} else {
		accountAccessToken, err := js.account.AccessToken()
		if err != nil {
			return Ticket{}
		}
		accessToken = accountAccessToken.AccessToken
	}

	apiUrl := fmt.Sprintf(
		"cgi-bin/ticket/getticket?access_token=%s&type=%s",
		accessToken,
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
	cache, cacheKeyPrefix := js.account.Cache()
	cacheKey := cacheKeyPrefix + js.getTicketCacheKey(ticketType)

	_ = cache.Put(
		cacheKey,
		string(ticketByte),
		time.Second*time.Duration(result.Ticket.ExpiresIn-600), // 提前过期
	)

	return result.Ticket
}

func (js *Jssdk) getTicketCacheKey(ticketType string) string {
	appId := js.account.AppId()
	if js.account.IsOpenPlatform() {
		appId = js.account.ComponentAppId() + "." + appId
	}
	return fmt.Sprintf(cacheKeyTemplate, ticketType, appId)
}

// configSignature jssdk 配置签名
func (js *Jssdk) configSignature() signatureConfig {
	nonce := support.RandString(10)
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
