package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUri = "https://api.weixin.qq.com/"

// Get get请求
func Get(uri string) ([]byte, error) {
	response, err := http.Get(baseUri + uri)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

// Post post请求
func Post(uri, data, contentType string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	if contentType == "" {
		contentType = "application/json;charset=utf-8"
	}
	response, err := http.Post(baseUri+uri, contentType, body)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

// PostJson Post Json 请求
func PostJson(uri string, data interface{}) ([]byte, error) {
	jsonBuf := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(baseUri+uri, "application/json;charset=utf-8", jsonBuf)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// PostJSONWithRespContentType post json数据请求，且返回数据类型
func PostJSONWithRespContentType(uri string, obj interface{}) ([]byte, string, error) {
	jsonBuf := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return nil, "", err
	}

	response, err := http.Post(baseUri+uri, "application/json;charset=utf-8", jsonBuf)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")
	return responseData, contentType, err
}
