package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	return io.ReadAll(response.Body)
}

// GetWithRespContentType get请求，并返回content-type
func GetWithRespContentType(uri string) ([]byte, string, error) {
	response, err := http.Get(baseUri + uri)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	responseData, err := io.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")
	return responseData, contentType, err
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

	return io.ReadAll(response.Body)
}

// PostJSON Post Json 请求
func PostJSON(uri string, data interface{}) ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(baseUri+uri, "application/json;charset=utf-8", jsonBuffer)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}

// PostJSONWithRespContentType post json数据请求，且返回数据类型
func PostJSONWithRespContentType(uri string, obj interface{}) ([]byte, string, error) {
	jsonBuffer := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return nil, "", err
	}

	response, err := http.Post(baseUri+uri, "application/json;charset=utf-8", jsonBuffer)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	responseData, err := io.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")
	return responseData, contentType, err
}
