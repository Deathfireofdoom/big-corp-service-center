package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/Deathfireofdoom/big-corp-service-center/the-manager/protobuf/themanager"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type TheManageServer struct {
	pb.UnimplementedTheManagerServer
}

func (s *TheManageServer) TestMessage(ctx context.Context, in *pb.TestMessage) (*pb.Message, error) {
	log.Printf("Receieved: %v", in.GetMessage())
	var random_string string = fmt.Sprintf("test-is-working-%v", rand.Intn(100))
	return &pb.Message{Message: in.GetMessage(), Number: in.GetNumber(), Random: random_string}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize new server
	s := grpc.NewServer()

	// Register server
	pb.RegisterTheManagerServer(s, &TheManageServer{})
	log.Printf("server listening at %v", lis.Addr())

	// start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
