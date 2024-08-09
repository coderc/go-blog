package login

import "github.com/gin-gonic/gin"

func LoginByDeviceHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login by device",
	})
}
