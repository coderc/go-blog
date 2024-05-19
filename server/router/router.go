package router

import (
	"github.com/coderc/go-blog/server/common/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use()

	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(middleware.Cors())
	v1.Any("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
