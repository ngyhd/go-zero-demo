package system

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/bff/internal/svc"
)

type SystemHandler struct {
	Version    string
	ServiceCtx *svc.ServiceContext
}

func NewSystemHandler(serviceCtx *svc.ServiceContext) *SystemHandler {
	return &SystemHandler{
		Version:    "dev",
		ServiceCtx: serviceCtx,
	}
}

func (h *SystemHandler) VersionHandler(w http.ResponseWriter, r *http.Request) {
	httpx.OkJson(w, map[string]string{
		"version": h.Version,
		"service": "bff",
	})
}

func (h *SystemHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	// 返回服务健康状态和时间戳
	// 详细健康检查（Redis、数据库、RPC服务）可以通过专门的监控接口实现
	httpx.OkJson(w, map[string]interface{}{
		"status":    "ok",
		"timestamp": time.Now().Unix(),
		"service":   "bff",
	})
}
