package rank

import (
	"context"
	"net/http"

	"github.com/coderc/go-blog/server/common/ret"
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
		ret.Response(c, http.StatusBadRequest, &ret.Res{
			Code:       http.StatusBadRequest,
			Message:    "invalid request",
			ErrMessage: err.Error(),
		})
		return
	}

	if rankItemSli, err := rank.GetRankList(context.TODO(), req.Start, req.End); err != nil {
		ret.Response(c, http.StatusInternalServerError, &ret.Res{
			Code:       http.StatusInternalServerError,
			Message:    "get rank list failed",
			ErrMessage: err.Error(),
		})
	} else {
		ret.Response(c, http.StatusOK, &ret.Res{
			Code:    http.StatusOK,
			Message: "success",
			Data:    rankItemSli,
		})
	}
}
