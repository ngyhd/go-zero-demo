package xerr

import (
	"fmt"
	"github.com/pkg/errors"
)

type BusinessError struct {
	Code int64
	Msg  string
}

func (e BusinessError) Error() string {
	return fmt.Sprintf("Business ErrCode:%d，ErrMsg:%s", e.Code, e.Msg)
}

func (e BusinessError) GetCode() int64 {
	return e.Code
}

func (e BusinessError) GetMessage() string {
	return e.Msg
}

func (e *BusinessError) SetMessage(msg string) BusinessError {
	e.Msg = msg
	return *e
}

func IsBusinessError(err error) bool {
	var e BusinessError
	if errors.Is(err, &e) {
		return true
	} else {
		return false
	}
}

func FromBusinessError(err error) BusinessError {
	var e BusinessError
	if errors.As(err, &e) {
		return e
	}
	return e
}

func New(code int64, msg string) BusinessError {
	return BusinessError{
		Code: code,
		Msg:  msg,
	}
}

// 通用错误
var (
	NotFoundErr = New(400000, "数据不存在")
)

// 用户服务错误
var (
	// 用户服务通用错误码
	AccountErr = New(400100, "业务错误")
)

// 推文服务错误
var (
	// 推文服务通用错误码
	PostErr = New(400200, "业务错误")
)
