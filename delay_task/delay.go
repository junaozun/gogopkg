package delay_task

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// Msg 消息
type Msg struct {
	Topic     string        // 消息的主题
	Key       string        // 消息的Key
	Body      []byte        // 消息的Body
	Delay     time.Duration // 延迟时间（秒）
	ReadyTime time.Time     // 消息准备好执行的时间（now + delay）
}

type PartitionRedisDelayQueue struct {
	client     *redis.Client // Redis客户端
	pushScript *redis.Script // Push脚本
	delScript  *redis.Script // Del脚本
}

func NewRedisDelayQueue(client *redis.Client) *PartitionRedisDelayQueue {
	return &PartitionRedisDelayQueue{
		client:     client,
		pushScript: redis.NewScript(delayQueuePushRedisScript),
		delScript:  redis.NewScript(delayQueueDelRedisScript),
	}
}

func (q *PartitionRedisDelayQueue) Push(ctx context.Context, msg *Msg) error {
	// 如果设置了ReadyTime，就使用RedisTime
	var readyTime int64
	if !msg.ReadyTime.IsZero() {
		readyTime = msg.ReadyTime.Unix()
	} else {
		// 否则使用Delay
		readyTime = time.Now().Add(msg.Delay).Unix()
	}
	success, err := q.pushScript.Run(ctx, q.client,
		[]string{q.topicZSet(msg.Topic), q.topicHash(msg.Topic)},
		msg.Key, msg.Body, readyTime).Bool()
	if err != nil {
		return err
	}
	if !success {
		return ErrDuplicateMessage
	}
	return nil
}

func (q *PartitionRedisDelayQueue) Delete(ctx context.Context, topic string, key string) error {
	return q.delScript.Run(ctx, q.client, []string{q.topicZSet(topic), q.topicHash(topic)}, key).Err()
}

var ErrDuplicateMessage = errors.New("duplicate message")

func (q *PartitionRedisDelayQueue) Consume(topic string, batchSize int, fn func(msg *Msg) error) {
	for {
		// 批量获取已经准备好执行的消息
		now := time.Now().Unix()
		zs, err := q.client.ZRangeByScoreWithScores(context.Background(), q.topicZSet(topic),
			&redis.ZRangeBy{
				Min:   "-inf",
				Max:   strconv.Itoa(int(now)),
				Count: int64(batchSize),
			}).Result()
		// 如果获取出错或者获取不到消息，则休眠一秒
		if err != nil || len(zs) == 0 {
			time.Sleep(time.Second)
			continue
		}
		// 遍历每个消息
		for _, z := range zs {
			key := z.Member.(string)
			// 获取消息的body
			body, err := q.client.HGet(context.Background(), q.topicHash(topic), key).Bytes()
			if err != nil {
				continue
			}

			// 处理消息
			err = fn(&Msg{
				Topic:     topic,
				Key:       key,
				Body:      body,
				ReadyTime: time.Unix(int64(z.Score), 0),
			})
			if err != nil {
				continue
			}

			// 如果消息处理成功，删除消息
			q.delScript.Run(context.Background(), q.client, []string{q.topicZSet(topic), q.topicHash(topic)}, key)
		}
	}
}

func (q *PartitionRedisDelayQueue) topicZSet(topic string) string {
	return fmt.Sprintf("%s:z", topic)
}

func (q *PartitionRedisDelayQueue) topicHash(topic string) string {
	return fmt.Sprintf("%s:h", topic)
}
