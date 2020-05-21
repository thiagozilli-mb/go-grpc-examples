package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"go-grpc-examples/unary/sum/sumpb"

	"google.golang.org/grpc"
)

var (
	val1 = flag.Int64("v1", 1, "Value A")
	val2 = flag.Int64("v2", 1, "Value B")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	c := sumpb.NewSumClient(conn)

	// numbers to add
	num := sumpb.Numbers{
		A: *val1,
		B: *val2,
	}

	// call Add service
	res, err := c.Add(context.Background(), &sumpb.SumRequest{Numbers: &num})
	if err != nil {
		log.Fatalf("failed to call Add: %v", err)
	}
	fmt.Println(res.Result)
}
