package error

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ApiError 微信Api返回的通用错误
type ApiError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Error 统一错误
type Error struct {
	Code int64 `json:"code"`
	Err  error `json:"error"`
}

func New(code int64, err error) error {
	return Error{Code: code, Err: err}
}

func (e Error) Error() string {
	return fmt.Sprintf("error: errcode=%d, errmsg=%s.", e.Code, e.Err.Error())
}

// Convert 错误转换
func Convert(err error) Error {
	codeRegxPatter := `errcode=(.*?),`
	msgRegxPatter := `errmsg=(.*?)\.`

	e := err.Error()

	codeMatcher := regexp.MustCompile(codeRegxPatter)
	code := codeMatcher.Find([]byte(e))
	codeString := strings.TrimRight(strings.Replace(string(code), "errcode=", "", -1), ",")
	codeInt, codeErr := strconv.ParseInt(codeString, 10, 64)
	if codeErr != nil {
		return Error{Code: 0, Err: err}
	}

	msgMatcher := regexp.MustCompile(msgRegxPatter)
	msg := msgMatcher.Find([]byte(e))
	msgString := strings.TrimRight(strings.Replace(string(msg), "errmsg=", "", -1), ".")
	if msgString == "" {
		return Error{Code: codeInt, Err: err}
	}

	return Error{
		Code: codeInt,
		Err:  errors.New(msgString),
	}
}
