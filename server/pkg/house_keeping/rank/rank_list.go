package rank

import (
	"context"
	"fmt"

	"github.com/coderc/go-blog/server/pkg/cache/redis"
	"github.com/pkg/errors"
	goRedis "github.com/redis/go-redis/v9"
)

func GetRankList(ctx context.Context, start, stop int64) ([]*CommitInfo, error) {
	var (
		deviceIdSli []string

		err error
	)

	if deviceIdSli, err = getDeviceIdSli(ctx, start, stop); err != nil {
		return nil, errors.WithMessage(err, "get device id slice failed")
	}

	return GetRankInfoByDeviceIdSli(ctx, deviceIdSli)
}

func getDeviceIdSli(ctx context.Context, start, stop int64) ([]string, error) {
	var (
		rankItemSli []goRedis.Z
		redisKey    = redis.GetPrefixHouseKeeping(redisKeyCommitScorePrefix)

		err error
	)

	if rankItemSli, err = redis.Client().ZRevRangeWithScores(ctx, redisKey, start, stop).Result(); err != nil {
		return nil, errors.WithMessage(err, "redis zrevrange withscores get rank list failed")
	}

	var deviceIdSli = make([]string, 0, len(rankItemSli))
	for _, item := range rankItemSli {
		deviceId, ok := item.Member.(string)
		if !ok {
			return nil, fmt.Errorf("rank item member type assert to string failed, rank item: %v", item)
		}
		deviceIdSli = append(deviceIdSli, deviceId)
	}

	return deviceIdSli, nil
}

func GetRankInfoByDeviceIdSli(ctx context.Context, deviceIdSli []string) ([]*CommitInfo, error) {
	var (
		redisKeySli = make([]string, 0, len(deviceIdSli))
	)

	for _, deviceId := range deviceIdSli {
		redisKeySli = append(redisKeySli, redis.GetPrefixHouseKeeping(redisKeyCommitInfoPrefix, deviceId))
	}

	pipe := redis.Client().Pipeline()

	var commitInfoCmdSli = make([]*goRedis.MapStringStringCmd, 0, len(redisKeySli))
	for _, redisKey := range redisKeySli {
		commitInfoCmdSli = append(commitInfoCmdSli, pipe.HGetAll(ctx, redisKey))
	}

	if cmdErrSli, err := pipe.Exec(ctx); err != nil {
		return nil, errors.WithMessage(err, "redis pipeline exec failed")
	} else {
		for _, cmdErr := range cmdErrSli {
			if cmdErr.Err() != nil {
				return nil, errors.WithMessage(cmdErr.Err(), "redis pipeline exec failed")
			}
		}
	}

	var commitInfoSli = make([]*CommitInfo, 0, len(commitInfoCmdSli))
	for _, commitInfoCmd := range commitInfoCmdSli {
		commitInfo := &CommitInfo{}
		if err := commitInfo.ConvertFromMap(commitInfoCmd.Val()); err != nil {
			return nil, errors.WithMessage(err, "commit info from redis map string string failed")
		}
		commitInfoSli = append(commitInfoSli, commitInfo)
	}

	return commitInfoSli, nil
}
