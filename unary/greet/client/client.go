package main

import (
	"context"
	"flag"
	"fmt"
	"go-grpc-examples/unary/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

var (
	name  = flag.String("first", "Thiago", "First Name")
	lname = flag.String("last", "Sarmento", "Last Name")
)

func main() {

	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	// create request
	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: *name,
			LastName:  *lname,
		},
	}

	// call Greet service
	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	fmt.Println(res.Result)
}
