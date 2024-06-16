package prome

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
)

var (
	// httpRequestCounter 记录每个路由的请求总数
	httpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Number of HTTP requests",
	}, []string{"path", "method", "httpStatus"})
)

func init() {
	// 注册指标对象
	prometheus.MustRegister(httpRequestCounter)
}

// Handle prome handler to gin handler
func Handle(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// Middleware prome 中间件 调用 prometheus 的指标对象
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		path := c.FullPath()
		method := c.Request.Method
		httpStatus := c.Writer.Status()
		httpRequestCounter.WithLabelValues(path, method, strconv.Itoa(httpStatus)).Inc()
	}
}
