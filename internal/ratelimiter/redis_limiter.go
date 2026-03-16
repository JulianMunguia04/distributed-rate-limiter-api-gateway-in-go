package ratelimiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func AllowRequest(ip string, limit int, window time.Duration) (bool, error) {

	key := "rate_limit:" + ip

	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		rdb.Expire(ctx, key, window)
	}

	if count > int64(limit) {
		return false, nil
	}

	return true, nil
}
