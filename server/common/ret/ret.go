package ret

import "github.com/gin-gonic/gin"

type Res struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	ErrMessage string      `json:"err_message"`
	Data       interface{} `json:"data"`
}

func Response(c *gin.Context, code int, res *Res) {
	c.JSON(code, res)
}
