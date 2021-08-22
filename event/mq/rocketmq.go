package mq

import (
	"context"
	"fmt"

	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2/consumer"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2/primitive"
	"git.n.xiaomi.com/streaming/rocketmq-client-go/v2/producer"
)

type Config struct {
	Nameserver []string `mapstructure:"name_server"`
	RetryTimes int      `mapstructure:"retry_times"`
	AK         string   `mapstructure:"ak"`
	SK         string   `mapstructure:"sk"`
	GroupName  string   `mapstructure:"group_name"`
}

type ProducerConfig struct {
	Config
}

type ConsumerConfig struct {
	Config
}

type RocketStore struct {
	producer     rocketmq.Producer
	PushConsumer rocketmq.PushConsumer
}

func NewRocketProducer(cfg ProducerConfig) (*RocketStore, error) {
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
	return &RocketStore{
		producer: p,
	}, nil
}

func NewRocketConsumer(cfg ConsumerConfig) (*RocketStore, error) {
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
	return &RocketStore{
		PushConsumer: c,
	}, nil
}

func (r *RocketStore) Close() error {
	if err := r.producer.Shutdown(); err != nil {
		return err
	}
	return nil
}

func (r *RocketStore) Create() error {
	return nil
}

func (r *RocketStore) Publish(ctx context.Context, topic, message string) error {
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

func (r *RocketStore) Subscribe(ctx context.Context, topic string) error {
	r.PushConsumer.Subscribe(
		topic,
		consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Printf("subscribe callback: %v \n", msgs[i])
			}
			return consumer.ConsumeSuccess, nil
		})
	return nil
}
