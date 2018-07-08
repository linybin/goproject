package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "github.com/linybin/goproject/protos/order"

	"fmt"
	"google.golang.org/grpc/reflection"
)

//models

type server struct {
}

func (s *server) GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.Order, error) {
	fmt.Println("we want to get order")
	order_id := request.Id
	return &pb.Order{		UserId: "2323",
		Id:     order_id,
		From:   "Hong Kong",
		To:     "China",
		Long:   32,
		Lat:    323,
		Status: "pending",
	}, nil
}

func (s *server) CreateOrder(ctx context.Context, request *pb.Order) (*pb.CreateOrderResponse, error) {
	return &pb.CreateOrderResponse{Success: true}, nil
}

func main() {
	fmt.Println("start the order grpc server ")
	//grpc
	lis, err := net.Listen("tcp", ":22222")
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
