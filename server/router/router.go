package router

import (
	"github.com/coderc/go-blog/server/common/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Any("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
