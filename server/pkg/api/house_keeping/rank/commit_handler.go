package rank

import (
	"context"
	"net/http"

	"github.com/coderc/go-blog/server/common/ret"
	"github.com/coderc/go-blog/server/pkg/house_keeping/rank"
	"github.com/gin-gonic/gin"
)

type CommitRequest struct {
	DeviceId string `json:"device_id" binding:"required"`
	Score    int64  `json:"score" binding:"required,min=0"`
}

func CommitScoreHandler(c *gin.Context) {
	var (
		req CommitRequest

		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		ret.Response(c, http.StatusBadRequest, &ret.Res{
			Code:       http.StatusBadRequest,
			Message:    "failed to parse request",
			ErrMessage: err.Error(),
		})
		return
	}

	if err = rank.CommitOneScore(context.TODO(), &rank.CommitScore{DeviceId: req.DeviceId, Score: req.Score}); err != nil {
		ret.Response(c, http.StatusInternalServerError, &ret.Res{
			Code:       http.StatusInternalServerError,
			Message:    "commit score failed",
			ErrMessage: err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	ret.Response(c, http.StatusOK, &ret.Res{
		Code:    http.StatusOK,
		Message: "success",
	})
}

type CommitInfoRequest struct {
	DeviceId string `json:"device_id" binding:"required"`
	Score    int64  `json:"score" binding:"required,min=0"`
	ShowName string `json:"show_name" binding:"required"`
}

func CommitInfoHandler(c *gin.Context) {
	var (
		req CommitInfoRequest

		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ret.Response(c, http.StatusBadRequest, &ret.Res{
			Code:       http.StatusBadRequest,
			Message:    "failed to parse request",
			ErrMessage: err.Error(),
		})
		return
	}

	if err = rank.CommitOneInfo(context.TODO(), &rank.CommitInfo{
		DeviceId: req.DeviceId,
		ShowName: req.ShowName,
		Score:    req.Score,
	}); err != nil {
		ret.Response(c, http.StatusInternalServerError, &ret.Res{
			Code:       http.StatusInternalServerError,
			Message:    "commit info failed",
			ErrMessage: err.Error(),
			Data:       nil,
		})
		return
	}

	ret.Response(c, http.StatusOK, &ret.Res{
		Code:    http.StatusOK,
		Message: "success",
	})
}
