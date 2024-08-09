package rank

import (
	"strconv"
	"time"

	"github.com/coderc/go-blog/server/pkg/cache/redis"
	"github.com/pkg/errors"
)

type CommitScore struct {
	DeviceId string `json:"device_id"`
	Score    int64  `json:"score"`
}

func (c *CommitScore) Key() (string, error) {
	if c == nil {
		return "", errors.New("commit is nil")
	}

	if c.DeviceId == "" {
		return "", errors.New("device id is empty")
	}

	return redis.GetPrefixHouseKeeping(redisKeyCommitScorePrefix), nil
}

func (c *CommitScore) CalcScore() float64 {
	a := float64(c.Score)
	b := float64(time.Now().UTC().Unix())
	for a < b {
		a *= 2
	}
	a /= b
	// 先提交的人排在前面
	return float64(c.Score) + a
}

type CommitInfo struct {
	DeviceId string `json:"device_id"`
	ShowName string `json:"show_name"`
	Score    int64  `json:"score"`
}

func (c *CommitInfo) Key() (string, error) {
	if c == nil {
		return "", errors.New("commit is nil")
	}

	if c.DeviceId == "" {
		return "", errors.New("device id is empty")
	}

	return redis.GetPrefixHouseKeeping(redisKeyCommitInfoPrefix, c.DeviceId), nil
}

func (c *CommitInfo) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"show_name": c.ShowName,
		"score":     c.Score,
		"device_id": c.DeviceId,
	}
}

func (c *CommitInfo) ConvertFromMap(hMap map[string]string) error {
	if hMap == nil {
		return errors.New("map is nil")
	}

	c.ShowName = hMap["show_name"]
	c.DeviceId = hMap["device_id"]
	var err error
	if c.Score, err = strconv.ParseInt(hMap["score"], 10, 64); err != nil {
		return errors.WithMessage(err, "parse score failed")
	}
	return nil
}

type RankItem struct {
	CommitScore
	CommitInfo
}
