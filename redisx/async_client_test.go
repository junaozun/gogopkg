package redisx

import (
	"context"
	"fmt"
	"testing"
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

func TestASyncClient_Del(t *testing.T) {
	key := "test"
	asyncClient.Del(key)
	asyncClient.Stop(context.Background())
}
