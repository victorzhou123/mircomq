package mq

import (
	"sync"

	"github.com/victorzhou123/simplemq/event"
)

type MQ interface {
	Push(*event.Message)
	Pop() event.Message
	HasMsg() bool
}

type mq struct {
	queue
}

func NewMQ() MQ {
	return &mq{
		queue: queue{
			queue: make([]*event.Message, 0),
		},
	}
}

type queue struct {
	m     sync.Mutex
	queue []*event.Message
}

func (q *queue) Push(msg *event.Message) {

	q.m.Lock()
	defer q.m.Unlock()

	q.queue = append(q.queue, msg)
}

func (q *queue) Pop() event.Message {

	q.m.Lock()
	defer q.m.Unlock()

	front := q.front()
	q.queue = q.queue[1:]

	return *front
}

func (q *queue) HasMsg() bool {
	return q.front() != nil
}

func (q *queue) front() *event.Message {
	if len(q.queue) > 0 {
		return q.queue[0]
	}

	return nil
}
