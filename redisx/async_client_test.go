package redisx

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestASyncClient_HMSet(t *testing.T) {
	data := map[string]interface{}{
		"name": "anhui",
		"age":  88,
	}
	key := "test"
	fields := NewArrayReq(4)
	for k, v := range data {
		if err := fields.Add(k, v); err != nil {
			t.Error(err)
		}
	}
	asyncClient.HMSet(key, fields)
	asyncClient.Stop(context.Background())
}

func TestASyncClient_HGet(t *testing.T) {
	key := "test"
	fields := []string{"name", "age"}
	asyncClient.HMGet(key, fields, func(strings []string, err error) {
		if err != nil {
			t.Error(err)
		}
		fmt.Println(strings)
	})
	asyncClient.Stop(context.Background())
}

func TestAsyncClient_ZRevRange(t *testing.T) {
	asyncClient.ZRevRange("su", 0, -1, func(res []redis.Z, err error) {
		if err != nil {
			t.Error(err)
		}
		fmt.Println(res)
	})
}

func TestASyncClient_Del(t *testing.T) {
	key := "test"
	asyncClient.Del(key)
	asyncClient.Stop(context.Background())
}
