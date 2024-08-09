package rank

import (
	"context"
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = rank.CommitOneScore(context.TODO(), &rank.CommitScore{DeviceId: req.DeviceId, Score: req.Score}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
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
		return
	}

	if err = rank.CommitOneInfo(context.TODO(), &rank.CommitInfo{
		DeviceId: req.DeviceId,
		ShowName: req.ShowName,
		Score:    req.Score,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
