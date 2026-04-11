package handler

import (
	"net/http"

	"go-zero-demo/pkg/upload"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type UploadHandler struct {
	storage *upload.LocalStorage
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{
		storage: upload.NewLocalStorage("./uploads", "/uploads"),
	}
}

func (h *UploadHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, err := upload.ParseFormFile(r, "file")
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	result, err := h.storage.Upload(r.Context(), file)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	httpx.OkJson(w, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}
