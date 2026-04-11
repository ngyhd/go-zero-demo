package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type SearchHandler struct{}

func NewSearchHandler() *SearchHandler {
	return &SearchHandler{}
}

func (h *SearchHandler) SearchPostsHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		httpx.ErrorCtx(r.Context(), w, &response{Code: 400, Message: "keyword is required"})
		return
	}
	// TODO: 调用 PostRpc 搜索推文
	httpx.OkJson(w, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data":    []interface{}{},
	})
}

func (h *SearchHandler) SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		httpx.ErrorCtx(r.Context(), w, &response{Code: 400, Message: "keyword is required"})
		return
	}
	// TODO: 调用 UserRpc 搜索用户
	httpx.OkJson(w, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data":    []interface{}{},
	})
}

type response struct {
	Code    int64
	Message string
}

func (e *response) Error() string {
	return e.Message
}
