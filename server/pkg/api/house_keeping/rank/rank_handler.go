package rank

import (
	"context"
	"net/http"

	"github.com/coderc/go-blog/server/pkg/house_keeping/rank"
	"github.com/gin-gonic/gin"
)

type RankListRequest struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

func RankListHandler(c *gin.Context) {
	var (
		req RankListRequest

		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	if rankItemSli, err := rank.GetRankList(context.TODO(), req.Start, req.End); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": rankItemSli,
		})
	}
}
