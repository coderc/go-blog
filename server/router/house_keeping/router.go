package house_keeping

import (
	"github.com/coderc/go-blog/server/pkg/api/house_keeping/rank"
	"github.com/coderc/go-blog/server/pkg/api/house_keeping/user/login"
	"github.com/gin-gonic/gin"
)

// InitHouseKeepingRouter 初始化house keeping 路由信息
func InitHouseKeepingRouter(r *gin.RouterGroup) {
	hGroup := r.Group("housekeeping")

	userApi := hGroup.Group("user")
	{
		loginApi := userApi.Group("login")
		loginApi.POST("login", login.LoginByDeviceHandler)
	}

	rankApi := hGroup.Group("rank")
	{
		rankApi.POST("commit/score", rank.CommitScoreHandler)
		rankApi.POST("commit/info", rank.CommitInfoHandler)
		rankApi.POST("list", rank.RankListHandler)
	}
}
