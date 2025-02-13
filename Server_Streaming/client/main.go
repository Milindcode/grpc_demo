package main

import (
	"context"
	"io"
	"log"

	sample "github.com/Milindcode/grpc_server_streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error connecting: %v\n", err)
	}

	client := sample.NewStockServiceClient(conn)

	stream, err := client.TestFunc(context.Background(), &sample.StockRequest{
		Symbol: "GOOGLE",
	})
	if err != nil {
		log.Printf("Error getting stream: %v\n", err)
	}

	for {
		data, err := stream.Recv()

		if err == io.EOF {
			log.Println("Stream closed.")
			break
		}
		if err != nil {
			log.Printf("Error recieving tick data: %v\n", err)
		}

		log.Println(data)
	}
}
