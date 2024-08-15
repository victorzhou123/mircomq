package publish

import (
	"github.com/victorzhou123/simplemq/event"
	"github.com/victorzhou123/simplemq/mq"
)

type Publisher interface {
	Publish(event event.Event)
}

type publishImpl struct {
	mq mq.MQ
}

func NewPublish(mq mq.MQ) publishImpl {
	return publishImpl{
		mq: mq,
	}
}

func (p *publishImpl) Publish(event event.Event) {

	if event.Message() == nil {
		return
	}

	p.mq.Push(event.Message())
}
