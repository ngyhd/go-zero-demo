package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Response 统一响应结构
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 返回成功响应
func Success(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, &Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// SuccessWithMessage 返回成功响应（带自定义消息）
func SuccessWithMessage(w http.ResponseWriter, message string, data interface{}) {
	httpx.OkJson(w, &Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

// Error 返回错误响应
func Error(w http.ResponseWriter, code int64, message string) {
	httpx.WriteJson(w, http.StatusOK, &Response{
		Code:    code,
		Message: message,
	})
}

// BadRequest 返回 400 错误
func BadRequest(w http.ResponseWriter, message string) {
	httpx.WriteJson(w, http.StatusBadRequest, &Response{
		Code:    400,
		Message: message,
	})
}

// Unauthorized 返回 401 错误
func Unauthorized(w http.ResponseWriter, message string) {
	httpx.WriteJson(w, http.StatusUnauthorized, &Response{
		Code:    401,
		Message: message,
	})
}

// Forbidden 返回 403 错误
func Forbidden(w http.ResponseWriter, message string) {
	httpx.WriteJson(w, http.StatusForbidden, &Response{
		Code:    403,
		Message: message,
	})
}

// NotFound 返回 404 错误
func NotFound(w http.ResponseWriter, message string) {
	httpx.WriteJson(w, http.StatusNotFound, &Response{
		Code:    404,
		Message: message,
	})
}

// InternalServerError 返回 500 错误
func InternalServerError(w http.ResponseWriter, message string) {
	httpx.WriteJson(w, http.StatusInternalServerError, &Response{
		Code:    500,
		Message: message,
	})
}
