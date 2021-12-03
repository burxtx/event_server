package mq

import (
	"context"
	"fmt"

	"git.n.xiaomi.com/op-basic/event_server/msg"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2/consumer"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2/primitive"
)

type Consumer struct {
	PushConsumer rocketmq.PushConsumer
	rocketmq.PullConsumer
}

func NewConsumer(cfg ConsumerConfig) (*Consumer, error) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(cfg.Nameserver),
		consumer.WithRetry(cfg.RetryTimes),
	)
	if err != nil {
		return nil, err
	}

	err = c.Start()
	if err != nil {
		return nil, err
	}
	return &Consumer{
		PushConsumer: c,
	}, nil
}

func (r *Consumer) Listen(ctx context.Context, topic string, subscription msg.ReceiveMessageFunc) error {
	return r.PushConsumer.Subscribe(
		topic,
		consumer.MessageSelector{},
		r.consumeMessage(ctx, subscription))
}

func (r *Consumer) Close() error {
	if err := r.PushConsumer.Shutdown(); err != nil {
		return err
	}
	return nil
}

func (r *Consumer) consumeMessage(ctx context.Context, receiver func(context.Context, msg.Message) error) func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	return func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	}

}
