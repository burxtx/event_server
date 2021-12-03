package mq

import (
	"context"

	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2/primitive"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2/producer"
)

type Producer struct {
	producer rocketmq.Producer
}

func NewProducer(cfg ProducerConfig) (*Producer, error) {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(cfg.Nameserver),
		producer.WithRetry(cfg.RetryTimes),
		producer.WithCredentials(primitive.Credentials{AccessKey: cfg.AK, SecretKey: cfg.SK}),
	)
	if err != nil {
		return nil, err
	}

	err = p.Start()
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: p,
	}, nil
}

func (r *Producer) Close() error {
	if err := r.producer.Shutdown(); err != nil {
		return err
	}
	return nil
}

func (r *Producer) Publish(ctx context.Context, topic, message string) error {
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(message),
	}
	_, err := r.producer.SendSync(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}
