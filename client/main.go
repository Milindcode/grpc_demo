package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Milindcode/grpc_demo/proto/samplepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting : %v", err)
	}
	defer conn.Close()

	client := samplepb.NewGreeterClient(conn)

	req := &samplepb.HelloRequest{
		Name: "Milind",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response : %v", err)
	}

	fmt.Println("Server Response:", res.Greeting)
}
