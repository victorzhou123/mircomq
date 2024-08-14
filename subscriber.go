package main

import "sync"

// Subscriber is a convenience return type for the Subscribe method.
type Subscriber interface {
	Consume(Event) error
	Topics() []string
}

type SubsMap map[string][]Subscriber

func (s SubsMap) Add(sub Subscriber) {

	for _, topic := range sub.Topics() {

		_, ok := s[topic]
		if !ok {
			s[topic] = []Subscriber{sub}

			continue
		}

		s[topic] = append(s[topic], sub)
	}
}

func (s SubsMap) Find(topic string) ([]Subscriber, bool) {

	arr, ok := s[topic]
	if !ok {
		return nil, false
	}

	if len(arr) == 0 {
		return nil, false
	}

	return arr, true
}

type SubscribeImpl interface {
	Handle(e Event)
}

type subscribeImpl struct {
	m    sync.Mutex
	subs SubsMap
}

func NewSubscribe() subscribeImpl {
	return subscribeImpl{}
}

func (s *subscribeImpl) Subscribe(sub Subscriber) {

	s.m.Lock()
	defer s.m.Unlock()

	s.subs.Add(sub)
}

func (s *subscribeImpl) Handle(e Event) {

	subs, ok := s.subs.Find(e.Topic())
	if !ok {
		return
	}

	for _, sub := range subs {
		sub.Consume(e)
	}
}
