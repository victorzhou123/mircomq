package main

type Publisher interface {
	Publish(event Event)
}

type publishImpl struct {
	mq MQ
}

func NewPublish(mq MQ) publishImpl {
	return publishImpl{
		mq: mq,
	}
}

func (p *publishImpl) Publish(event Event) {

	if event.Message() == nil {
		return
	}

	p.mq.Push(event.Message())
}
