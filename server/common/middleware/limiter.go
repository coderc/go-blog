package middleware

import (
	"net/http"

	"github.com/coderc/go-blog/server/common/ret"
	"github.com/coderc/go-blog/server/pkg/house_keeping/limiter"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if incr, err := limiter.Incr(c.ClientIP()); err != nil {
			ret.Response(c, http.StatusInternalServerError, &ret.Res{
				Code:       http.StatusInternalServerError,
				Message:    "limiter failed",
				ErrMessage: err.Error(),
			})
			c.Abort()
		} else if incr > 50 || c.GetHeader("foo") != "bar" {
			ret.Response(c, http.StatusForbidden, &ret.Res{
				Code:       http.StatusForbidden,
				Message:    "limiter failed",
				ErrMessage: "too many requests",
			})
			c.Abort()
		}
		c.Next()
	}
}
