package etcdx

import (
	"context"
	"testing"
)

func TestPut(t *testing.T) {
	cfg := Config{
		Servers: "10.18.98.163:2379,10.18.98.164:2379,10.18.98.165:2379",
	}
	cli, err := NewClientWithConfig(cfg)
	if err != nil {
		t.Error(err)
	}
	err = cli.Put(context.Background(), "/nihao/", "nihao")
	if err != nil {
		t.Error(err)
	}
	cli.Close(context.Background())
}
