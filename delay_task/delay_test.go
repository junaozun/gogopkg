package delay_task

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestNewPartitionRedisDelayQueue(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
	})
	q := NewRedisDelayQueue(client)
	topic := "{wenwen}_test"
	go q.Consume(topic, 5, func(msg *Msg) error {
		fmt.Printf("consume partiton0: %+v\n", msg)
		return nil
	})
	for {

	}
}

func TestPush(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "1234567",
	})
	q := NewRedisDelayQueue(client)
	count := 10
	for i := 0; i < count; i++ {
		err := q.Push(context.Background(), &Msg{
			Topic: "{wenwen}_test",
			Key:   "ss" + strconv.Itoa(i),
			Body:  []byte("test" + strconv.Itoa(i)),
			Delay: time.Second * time.Duration(30),
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}

func TestDelete(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	q := NewRedisDelayQueue(client)
	topic := "test"
	err := q.Delete(context.Background(), topic, "suxuefeng0")
	if err != nil {
		t.Error(err)
	}
}
