package main

import (
	"context"
	"fmt"
	"grpc_project/multiplypb/multiplypb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	port := "localhost:8000"
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("connection failed at : %v", err)
	}

	defer conn.Close()

	newUser := multiplypb.NewMultiplyServiceClient(conn)

	numbers := multiplypb.Numbers{
		A: 5,
		B: 4,
	}

	result, err := newUser.Multiply(context.Background(), &multiplypb.MultiplyRequest{Numbers: &numbers})

	if err != nil {
		log.Fatalf("service failed at: %v\n", err)
	}
	fmt.Println("the result is:", result)
}
