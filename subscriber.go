package main

import "sync"

type Consumer interface {
	Consume(Event)
	Topics() []string
}

type SubsMap map[string][]Consumer

func newEmptySubsMap() SubsMap {
	return make(map[string][]Consumer, 0)
}

func (s SubsMap) Add(sub Consumer) {

	if len(sub.Topics()) == 0 {
		return
	}

	for _, topic := range sub.Topics() {

		_, ok := s[topic]
		if !ok {
			s[topic] = []Consumer{sub}

			continue
		}

		s[topic] = append(s[topic], sub)
	}
}

func (s SubsMap) Find(topic string) ([]Consumer, bool) {

	arr, ok := s[topic]
	if !ok {
		return nil, false
	}

	if len(arr) == 0 {
		return nil, false
	}

	return arr, true
}

type Distributer interface {
	Distribute(e Event)
}

type subscribeImpl struct {
	m    sync.Mutex
	subs SubsMap
}

func NewSubscribeImpl() subscribeImpl {
	return subscribeImpl{
		subs: newEmptySubsMap(),
	}
}

func (s *subscribeImpl) Subscribe(sub Consumer) {

	s.m.Lock()
	defer s.m.Unlock()

	s.subs.Add(sub)
}

func (s *subscribeImpl) Distribute(e Event) {

	subs, ok := s.subs.Find(e.Topic())
	if !ok {
		return
	}

	for _, sub := range subs {
		sub.Consume(e)
	}
}
