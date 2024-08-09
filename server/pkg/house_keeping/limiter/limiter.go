package limiter

import (
	"context"
	"time"

	"github.com/coderc/go-blog/server/pkg/cache/redis"
	"github.com/pkg/errors"
)

const (
	redisKeyPrefix = "limiter"
)

func Incr(key string) (int64, error) {
	redisKey := redis.GetPrefixHouseKeeping(redisKeyPrefix, key, time.Now().UTC().Format("2006-01-02-15-04"))
	if incr, err := redis.Client().Incr(context.TODO(), redisKey).Result(); err != nil {
		return 0, errors.WithMessage(err, "redis.Incr failed")
	} else {
		return incr, nil
	}
}
