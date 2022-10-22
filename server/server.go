package main

import (
	"context"
	"grpc_project/multiplypb/multiplypb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Multiply(ctx context.Context, req *multiplypb.MultiplyRequest) (*multiplypb.MultiplyResponse, error) {
	a, b := req.GetNumbers().GetA(), req.GetNumbers().GetB()
	result := a * b
	return &multiplypb.MultiplyResponse{Result: result}, nil
}

func main() {
	port := "localhost:8000"
	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("listen error at: %v", err)
	}

	s := grpc.NewServer()
	multiplypb.RegisterMultiplyServiceServer(s, &server{})

	err = s.Serve(listen)

	if err != nil {
		log.Fatalf("server error at: %v", err)
	}
}
