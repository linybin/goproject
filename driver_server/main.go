package main

import (
	"log"
	"net"

	pb "benlin/helloworld"
	"google.golang.org/grpc"

	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) GetDriver(ctx context.Context, req *pb.GetDriverRequest) (*pb.Driver, error) {
	id := req.Id
	driver := pb.Driver{
		Name: "ben",
		Id:   id,
	}
	return &driver, nil
}

func main() {
	fmt.Println("Start the grpc server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDriverServerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
