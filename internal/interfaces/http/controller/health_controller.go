package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
)

func init() {
	di.AddSingleton[*HealthController](NewHealthController)
}

// HealthController 健康检查控制器
type HealthController struct {
	logger *log.Logger
}

// NewHealthController 创建健康检查控制器
func NewHealthController() (*HealthController, error) {
	return &HealthController{
		logger: di.Get[*log.Logger](),
	}, nil
}

// HealthCheckResponse 健康检查响应
type HealthCheckResponse struct {
	Status    string            `json:"status"`
	Version   string            `json:"version"`
	Services  map[string]string `json:"services"`
	Timestamp int64             `json:"timestamp"`
}

// HealthCheck 健康检查
func (c *HealthController) HealthCheck(ctx context.Context, reqCtx *app.RequestContext) {
	response := HealthCheckResponse{
		Status:  "healthy",
		Version: "1.0.0",
		Services: map[string]string{
			"database": "healthy",
			"cache":    "healthy",
		},
		Timestamp: time.Now().Unix(),
	}

	dto.Ok(reqCtx, "Service is healthy", response)
}

// ReadinessCheck 就绪检查
func (c *HealthController) ReadinessCheck(ctx context.Context, reqCtx *app.RequestContext) {
	// 这里可以检查依赖服务的可用性
	// 例如数据库连接、缓存连接等

	reqCtx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ready",
	})
}

// LivenessCheck 存活检查
func (c *HealthController) LivenessCheck(ctx context.Context, reqCtx *app.RequestContext) {
	reqCtx.JSON(http.StatusOK, map[string]interface{}{
		"status": "alive",
	})
}
