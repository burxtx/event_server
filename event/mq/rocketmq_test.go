package mq

import (
	"context"
	"testing"
)

var Nameserver = "http://staging-cnbj2-rocketmq.namesrv.api.xiaomi.net:9876"
var TOPIC = "cmdb-test"
var AK = ""
var SK = ""

func TestRocketProducer(t *testing.T) {
	cfg := Config{
		Nameserver: []string{Nameserver},
		AK:         AK,
		SK:         SK,
	}
	producer, err := NewRocketProducer(ProducerConfig{cfg})
	if err != nil {
		t.Errorf(err.Error())
	}
	err = producer.Publish(context.Background(), TOPIC, "hello rocketmq")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRocketConsumer(t *testing.T) {
	cfg := Config{Nameserver: []string{Nameserver},
		AK: AK,
		SK: SK,
	}
	producer, err := NewRocketConsumer(ConsumerConfig{cfg})
	if err != nil {
		t.Errorf(err.Error())
	}
	err = producer.Subscribe(context.Background(), TOPIC)
	if err != nil {
		t.Errorf(err.Error())
	}
}
