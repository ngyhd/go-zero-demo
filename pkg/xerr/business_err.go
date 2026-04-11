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
	NotFoundErr    = New(400000, "数据不存在")
	SystemErr      = New(400001, "系统错误")
	ParamErr       = New(400002, "参数错误")
	UnauthorizedErr = New(400003, "未授权")
)

// 用户服务错误
var (
	// 用户相关错误码 (40xxxx)
	AccountErr          = New(401000, "账号错误")
	AccountNotFoundErr  = New(401001, "账号不存在")
	AccountExistsErr    = New(401002, "账号已存在")
	PasswordErr         = New(401003, "密码错误")
	PasswordFormatErr   = New(401004, "密码格式不匹配")
	AccountFormatErr    = New(401005, "账号格式不匹配")
	NoPermissionErr    = New(401006, "无权限操作")
	UserDisabledErr     = New(401007, "账号已注销")
)

// 推文服务错误
var (
	// 推文相关错误码 (42xxxx)
	PostErr          = New(422000, "推文错误")
	PostNotFoundErr  = New(422001, "推文不存在")
	PostTitleErr     = New(422002, "推文标题不能为空")
	PostContentErr    = New(422003, "推文内容不能为空")
	NoPostPermissionErr = New(422004, "无权操作此推文")
)
