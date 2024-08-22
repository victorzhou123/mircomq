package internal

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/victorzhou123/simplemq/event"
	pb "github.com/victorzhou123/simplemq/event/message"
)

type server struct {
	pb.UnimplementedMqServer

	mq MQ
}

func NewServer(mq MQ) pb.MqServer {
	return &server{
		mq: mq,
	}
}

func (s *server) Pop(ctx context.Context, empty *emptypb.Empty) (*pb.Message, error) {

	msg := s.mq.Pop()

	return &pb.Message{Key: msg.MessageKey(), Body: msg.Body}, nil
}

func (s *server) Push(ctx context.Context, msg *pb.Message) (*emptypb.Empty, error) {

	eventMsg := &event.Message{}
	eventMsg.SetMessageKey(msg.Key)
	eventMsg.Body = msg.GetBody()

	s.mq.Push(eventMsg)

	return &emptypb.Empty{}, nil
}

func (s *server) HasMsg(ctx context.Context, empty *emptypb.Empty) (*pb.BoolMsg, error) {

	b := s.mq.HasMsg()

	return &pb.BoolMsg{Val: b}, nil
}
