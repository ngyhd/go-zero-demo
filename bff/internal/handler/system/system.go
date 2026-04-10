package system

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type SystemHandler struct {
	Version string
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{
		Version: "dev",
	}
}

func (h *SystemHandler) VersionHandler(w http.ResponseWriter, r *http.Request) {
	httpx.OkJson(w, map[string]string{
		"version": h.Version,
		"service": "bff",
	})
}

func (h *SystemHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	httpx.OkJson(w, map[string]string{"status": "ok"})
}
