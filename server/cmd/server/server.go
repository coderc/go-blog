package main

import (
	"fmt"

	"github.com/coderc/go-blog/server/pkg/cache/redis"
	"github.com/coderc/go-blog/server/pkg/config"
	"github.com/coderc/go-blog/server/pkg/logger"
	"github.com/coderc/go-blog/server/router"
	"github.com/gin-gonic/gin"
)

func Init() {
	config.Init()
	logger.Init()
	redis.Init()
}

func main() {
	Init()

	r := gin.New()
	router.Init(r)
	if err := r.Run(fmt.Sprintf(":%d", config.GetConfig().Server.Port)); err != nil {
		panic(err)
	}
}
