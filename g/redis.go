package g

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

type myRedis struct {
	*redis.Client
	once sync.Once
}

var rdb myRedis

func RDB() *redis.Client {
	rdb.once.Do(func() {
		rdb.Client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", VP.GetString("redis.host"), VP.GetInt("redis.port")),
			Password: VP.GetString("redis.password"),
			DB:       VP.GetInt("redis.db"),
		})
		var ctx = context.Background()
		if _, err := rdb.Ping(ctx).Result(); err != nil {
			panic(fmt.Errorf("无法连接到redis: %s \n", err))
		}
	})
	return rdb.Client
}
