package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"testing"
)

type As struct {
	Name string
	Age  int
}

func TestRedis(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       9,
	})

	marshal, _ := json.Marshal(As{Name: "QAQ", Age: 10})

	err := rdb.Set(ctx, "key",marshal , 0).Err()
	if err != nil {
		panic(err)
	}
	res := rdb.Ping(ctx)
	println(res.String())
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		println(errors.Is(err, redis.Nil))
		return
	}
	fmt.Println("key", val)
}
