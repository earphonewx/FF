package utils

import (
	"context"
	"errors"
	"ff/g"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedLock struct {
	Key   string
	Value string
	Expiry time.Duration
	Ctx    context.Context
}

func (r *RedLock) Lock() error {
	 if res, err := g.RDB().SetNX(r.Ctx, r.Key, r.Value, r.Expiry).Result(); !res || err != nil {
	 	return errors.New("获取redlock失败")
	 }
	 return nil
}

var deleteScript = redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then 
			return redis.call("del", KEYS[1]) 
		else 
			return 0
		end
`)

func (r *RedLock) Unlock() error {
	return deleteScript.Run(r.Ctx, g.RDB(), []string{r.Key}, r.Value).Err()
}
