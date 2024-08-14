package main

import (
	"sync"
)

type MQ interface {
	Push(*Message)
	Pop() Message
	HasMsg() bool
}

type mq struct {
	queue
}

func NewMQ() MQ {
	return &mq{
		queue: queue{
			queue: make([]*Message, 0),
		},
	}
}

type queue struct {
	m     sync.Mutex
	queue []*Message
}

func (q *queue) Push(msg *Message) {

	q.m.Lock()
	defer q.m.Unlock()

	q.queue = append(q.queue, msg)
}

func (q *queue) Pop() Message {

	q.m.Lock()
	defer q.m.Unlock()

	front := q.front()
	q.queue = q.queue[1:]

	return *front
}

func (q *queue) HasMsg() bool {
	return q.front() != nil
}

func (q *queue) front() *Message {
	if len(q.queue) > 0 {
		return q.queue[0]
	}

	return nil
}
