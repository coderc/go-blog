package redis

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/coderc/go-blog/server/pkg/config"
	"github.com/coderc/go-blog/server/pkg/logger"
	goRedis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	incrIndex atomic.Uint64
	clientMap map[uint64]*goRedis.Client
)

func Init() {
	incrIndex.Store(0)
	clientMap = make(map[uint64]*goRedis.Client)
	refreshClient()
}

func Client() *goRedis.Client {
	return clientMap[incrIndex.Load()]
}

func refreshClient() {
	ok := false
	for i := 0; i < 3; i++ {
		if client, err := getRedisClient(); err != nil {
			logger.B.Error("refresh redis client error", zap.Error(err))
		} else {
			ok = true
			clientMap[incrIndex.Add(1)] = client
			break
		}
	}

	if !ok {
		logger.B.Error("refresh redis client error after 3 times")
	} else {
		delete(clientMap, incrIndex.Load()-1) // 删除上一个 后续等待gc会回收
	}

}

func getRedisClient() (*goRedis.Client, error) {
	client := goRedis.NewClient(&goRedis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetConfig().Redis.Host, config.GetConfig().Redis.Port),
		Password: config.GetConfig().Redis.Password,
		DB:       config.GetConfig().Redis.Db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), config.GetConfig().Redis.PingTimeout)
	defer cancel()
	val, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	if val != "PONG" {
		return nil, errors.New("redis ping result is not PONG")
	}

	return client, nil
}

func GetPrefixHouseKeeping[T string | int | int32 | int64 | uint32 | uint | uint64](keySli ...T) string {
	b := strings.Builder{}
	b.WriteString(config.GetConfig().Redis.PrefixHouseKeeping)
	for _, key := range keySli {
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%v", key))
	}
	return b.String()
}
