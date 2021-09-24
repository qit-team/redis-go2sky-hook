package redis_go2sky_hook

import (
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	goredis "github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestNewSkyWalkingHook(t *testing.T) {
	tracer, err := go2sky.NewTracer("127.0.0.1:11800")
	if err != nil {
		t.Error(err)
		return
	}
	var ctx = context.Background()
	hook := NewSkyWalkingHook(tracer)
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, // use default DB
	})
	rdb.AddHook(hook)
	err = rdb.Set(ctx, "key", "test-hook", 100*time.Second).Err()
	if err != nil {
		t.Error(err)
		return
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("redis get result:%v", val)
}
