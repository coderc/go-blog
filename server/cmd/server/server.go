package main

import (
	"github.com/coderc/go-blog/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	router.Init(r)
	if err := r.Run(":8082"); err != nil {
		panic(err)
	}
}
