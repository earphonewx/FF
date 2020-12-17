package initialize

import (
	"ff/g"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	g.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", g.VP.GetString("redis.host"), g.VP.GetInt("redis.port")),
		Password: g.VP.GetString("redis.password"),
		DB:       g.VP.GetInt("redis.db"),
	})
	//get := g.Redis.Get(ctx, "key")
	fmt.Println(g.Redis)
}
