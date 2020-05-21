package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"go-grpc-examples/unary/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.Println("Server start on :50051...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}

// Greet greets with FirstName and LastMame
func (*server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet from %s %s, answer back message\n!",
		in.GetGreeting().GetFirstName(),
		in.GetGreeting().GetLastName())

	result := fmt.Sprintf("Hello %s %s (:",
		in.GetGreeting().GetFirstName(),
		in.GetGreeting().GetLastName())

	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}
