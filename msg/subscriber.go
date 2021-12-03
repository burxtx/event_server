package msg

import (
	"sync"

	"git.n.xiaomi.com/op-basic/event_server/libs/log"
)

type MessageSubscriber interface {
	Subscribe(channel string, receiver MessageReceiver)
}

// Subscriber receives domain events, commands, and replies from the consumer
type Subscriber struct {
	consumer     Consumer
	logger       log.MyLogger
	middlewares  []func(MessageReceiver) MessageReceiver
	receivers    map[string][]MessageReceiver
	stopping     chan struct{}
	subscriberWg sync.WaitGroup
	close        sync.Once
}

// NewSubscriber constructs a new Subscriber
func NewSubscriber(consumer Consumer) *Subscriber {
	s := &Subscriber{
		consumer:  consumer,
		receivers: make(map[string][]MessageReceiver),
		stopping:  make(chan struct{}),
		logger:    log.MyLogger,
	}

	s.logger.Info("msg.Subscriber constructed")

	return s
}

// Subscribe connects the receiver with messages from the channel on the consumer
func (s *Subscriber) Subscribe(channel string, receiver MessageReceiver) {
	if _, exists := s.receivers[channel]; !exists {
		s.receivers[channel] = []MessageReceiver{}
	}
	s.logger.Info("subscribed", "Channel", channel)
	s.receivers[channel] = append(s.receivers[channel], s.chain(receiver))
}

func (s *Subscriber) chain(receiver MessageReceiver) MessageReceiver {
	if len(s.middlewares) == 0 {
		return receiver
	}

	r := s.middlewares[len(s.middlewares)-1](receiver)
	for i := len(s.middlewares) - 2; i >= 0; i-- {
		r = s.middlewares[i](r)
	}

	return r
}

// Use appends middleware receivers to the receiver stack
func (s *Subscriber) Use(mws ...func(MessageReceiver) MessageReceiver) {
	if len(s.receivers) > 0 {
		panic("middleware must be added before any subscriptions are made")
	}

	s.middlewares = append(s.middlewares, mws...)
}
