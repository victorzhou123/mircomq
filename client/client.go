package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/victorzhou123/simplemq/event"
	pb "github.com/victorzhou123/simplemq/event/message"
	"github.com/victorzhou123/simplemq/internal"
)

type Client interface {
	internal.MQ
	Close()
}

type client struct {
	conn *grpc.ClientConn

	expire   time.Duration
	mqClient pb.MqClient
}

func NewClient(addr string, t time.Duration) (Client, error) {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &client{
		conn:     conn,
		mqClient: pb.NewMqClient(conn),
		expire:   t,
	}, nil
}

func (c *client) Push(msg *event.Message) {

	ctx, cancel := context.WithTimeout(context.Background(), c.expire)
	defer cancel()

	_, err := c.mqClient.Push(ctx, &pb.Message{
		Key:  msg.MessageKey(),
		Body: msg.Body,
	})
	if err != nil {
		log.Fatalf("push error: %s", err.Error())
	}
}

func (c *client) Pop() event.Message {

	ctx, cancel := context.WithTimeout(context.Background(), c.expire)
	defer cancel()

	msg, err := c.mqClient.Pop(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("pop error: %s", err.Error())
	}

	message := event.Message{}
	message.SetMessageKey(msg.GetKey())
	message.Body = msg.GetBody()

	return message
}

func (c *client) HasMsg() bool {

	ctx, cancel := context.WithTimeout(context.Background(), c.expire)
	defer cancel()

	bMsg, err := c.mqClient.HasMsg(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("hasMsg error: %s", err.Error())
	}

	return bMsg.Val
}

func (c *client) Close() {
	_ = c.conn.Close()
}
