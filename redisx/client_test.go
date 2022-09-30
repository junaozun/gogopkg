package redisx

import (
	"context"
	"os"
	"testing"
)

var (
	client      IClient
	asyncClient *AsyncClient
)

func TestMain(m *testing.M) {
	cfg := Config{
		Server: "127.0.0.1:6379",
		Index:  0,
	}
	var err error
	client, err = NewClient(cfg)
	if err != nil {
		panic(err)
	}
	pushFunc := func(ctx context.Context, f func()) error {
		// fmt.Println("pushFunc")
		f()
		return nil
	}
	asyncClient = NewAsync(client, WithPushFunc(pushFunc))
	os.Exit(m.Run())
}
