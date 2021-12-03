package msg

import (
	"sync"

	"git.n.xiaomi.com/op-basic/event_server/libs/log"
)

// Publisher send domain events, commands, and replies to the publisher
type Publisher struct {
	producer Producer
	logger   log.MyLogger
	close    sync.Once
}

// NewPublisher constructs a new Publisher
func NewPublisher(producer Producer) *Publisher {
	p := &Publisher{
		producer: producer,
		logger:   log.MyLogger,
	}

	p.logger.Info("msg.Publisher constructed")

	return p
}
