package rank

import (
	"context"

	"github.com/coderc/go-blog/server/pkg/cache/redis"
	"github.com/pkg/errors"
	goRedis "github.com/redis/go-redis/v9"
)

const (
	redisKeyCommitScorePrefix = "rank:commit:score"
	redisKeyCommitInfoPrefix  = "rank:commit:info"
)

// CommitOneScore 提交一个玩家的分数到排行榜
func CommitOneScore(ctx context.Context, commit *CommitScore) error {
	var (
		redisKey string

		err error
	)

	if redisKey, err = commit.Key(); err != nil {
		return errors.WithMessage(err, "get commit key failed")
	}

	var oldScore float64
	oldScore, _ = getScoreByDeviceId(ctx, commit.DeviceId)

	if oldScore > commit.CalcScore() {
		return nil
	}

	if err = redis.Client().ZAdd(ctx, redisKey, goRedis.Z{
		Score:  commit.CalcScore(),
		Member: commit.DeviceId,
	}).Err(); err != nil {
		return errors.WithMessage(err, "redis zadd failed")
	}

	return nil
}

func CommitOneInfo(ctx context.Context, commit *CommitInfo) error {
	var (
		redisKey string

		err error
	)

	if redisKey, err = commit.Key(); err != nil {
		return errors.WithMessage(err, "get commit info key failed")
	}

	if err = redis.Client().HSet(ctx, redisKey, commit.GetMap()).Err(); err != nil {
		return errors.WithMessage(err, "redis hset failed")
	}

	return nil
}

func getScoreByDeviceId(ctx context.Context, deviceId string) (float64, error) {
	redisKey := redis.GetPrefixHouseKeeping(redisKeyCommitScorePrefix)
	return redis.Client().ZScore(ctx, redisKey, deviceId).Result()
}
