package main

import (
	"bytes"
	"testing"
)

const MsgBodyThisIsMessageBody = "this is message body"

var (
	buf                                  = &bytes.Buffer{}
	noTopicConsumer, withTopicConsumer   Consumer
	noTopicConsumer2, withTopicConsumer2 Consumer
	missTopicMsg, hitTopicMsg            Message
)

type cs1 struct{}

func (s *cs1) Consume(e Event) {
	e.Message()
}

func (s *cs1) Topics() []string {
	return []string{"topic1", "topic2"}
}

type cs2 struct{}

func (s *cs2) Consume(e Event) {
	e.Message()
}

func (s *cs2) Topics() []string {
	return nil
}

func init() {
	// Test_subscribeImpl_Subscribe
	noTopicConsumer = &cs1{}
	withTopicConsumer = &cs2{}

	// Test_subscribeImpl_Handle
	noTopicConsumer2 = &cs3{buf}
	withTopicConsumer2 = &cs4{buf}
	missTopicMsg = Message{
		key:    "topic_not_hit",
		Header: make(map[string]string),
		Body:   []byte{},
	}
	hitTopicMsg = Message{
		key:    "topic4",
		Header: make(map[string]string),
		Body:   []byte(MsgBodyThisIsMessageBody),
	}
}

func Test_subscribeImpl_Subscribe(t *testing.T) {
	type args struct {
		sub Consumer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"no topics",
			args{
				noTopicConsumer,
			},
		},
		{
			"with topics return",
			args{
				withTopicConsumer,
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

type cs3 struct {
	buffer *bytes.Buffer
}

func (s *cs3) Consume(e Event) {
	s.buffer.Write(e.Message().Body)
}

func (s *cs3) Topics() []string {
	return nil
}

type cs4 struct {
	buffer *bytes.Buffer
}

func (s *cs4) Consume(e Event) {
	s.buffer.Write(e.Message().Body)
}

func (s *cs4) Topics() []string {
	return []string{"topic3", "topic4"}
}

func Test_subscribeImpl_Handle(t *testing.T) {
	type args struct {
		c Consumer
		e Event
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"miss topic",
			args{
				noTopicConsumer2,
				missTopicMsg,
			},
			"",
		},
		{
			"hit topic",
			args{
				withTopicConsumer2,
				hitTopicMsg,
			},
			MsgBodyThisIsMessageBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSubscribeImpl()
			s.Subscribe(tt.args.c)
			s.Distribute(tt.args.e)

			got := buf.String()

			if got != tt.want {
				t.Errorf("got %s want %s", got, tt.want)
			}
		})
	}
}
