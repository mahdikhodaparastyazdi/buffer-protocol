package greetserver

import (
	"fmt"

	"grpc_course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {

	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listion: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed: %v", err)
	}
}
