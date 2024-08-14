package main

import (
	"testing"
)

var s1, s2 Subscriber

type sub1 struct{}

func (s *sub1) Consume(e Event) {
	e.Message()
}

func (s *sub1) Topics() []string {
	return []string{"topic1", "topic2"}
}

type sub2 struct{}

func (s *sub2) Consume(e Event) {
	e.Message()
}

func (s *sub2) Topics() []string {
	return nil
}

func init() {
	s1 = &sub1{}
	s2 = &sub2{}
}

func Test_subscribeImpl_Subscribe(t *testing.T) {
	type args struct {
		sub Subscriber
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"no topics",
			args{
				s1,
			},
		},
		{
			"with topics return",
			args{
				s2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSubscribeImpl()
			s.Subscribe(tt.args.sub)
		})
	}
}
