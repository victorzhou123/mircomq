package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/victorzhou123/simplemq/event/message"
	"github.com/victorzhou123/simplemq/internal"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// mq init
	mq := internal.NewMQ()

	// mq server
	mqSvc := internal.NewServer(mq)

	s := grpc.NewServer()
	pb.RegisterMqServer(s, mqSvc)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
