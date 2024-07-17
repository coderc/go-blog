package router

import (
	"net/http"

	"github.com/coderc/go-blog/server/common/middleware"
	"github.com/coderc/go-blog/server/common/prome"
	"github.com/coderc/go-blog/server/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func Init(r *gin.Engine) {
	r.Use()

	r.Use(prome.Middleware())
	r.GET("/api/metrics", prome.Handle(promhttp.Handler()))

	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(middleware.Cors())
	v1.Any("/ping", func(c *gin.Context) {
		logger.B.Info("ping", zap.String("ip", c.ClientIP()))
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.Any("/mock_alert", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "mock alert",
		})
	})
}
