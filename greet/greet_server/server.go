package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listion: %v", err)
	}
	s := grpc.NewServer()
}
