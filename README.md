# redis-go2sky-hook

skywalking hook for redis

## Example
```go
package main

import (
    "log"
    "context"
    "time"

    redisSkyHook "github.com/qit-team/redis-go2sky-hook"
	"github.com/SkyAPM/go2sky"
	goredis "github.com/go-redis/redis/v8"
)


func main() {
	tracer, err := go2sky.NewTracer("127.0.0.1:11800")
	if err != nil {
		log.Fatal(err)
	}
	var ctx = context.Background()
	hook := redisSkyHook.NewSkyWalkingHook(tracer)
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, // use default DB
	})
	rdb.AddHook(hook)
	err = rdb.Set(ctx, "key", "test-hook", 100*time.Second).Err()
	...
}
```