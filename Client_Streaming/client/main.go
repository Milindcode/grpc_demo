package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	sample "github.com/Milindcode/grpc_client_streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting : %v", err)
	}
	defer conn.Close()

	client := sample.NewFileStreamerClient(conn)
	fileName := "Paper1.pdf"
	filePath := "data/" + fileName

	stream, err := client.SendFile(context.Background())
	if err != nil {
		log.Printf("Error getting stream: %v\n", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	buffer := make([]byte, 8192)
	count:= 1
	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error recieving chunk: %v", err)
		}

		req := &sample.FileData{
			Chunk:    buffer[:n],
			FileName: fileName,
		}

		log.Printf("Sending chunk-%v\n", count)
		count++ 

		if err := stream.Send(req); err != nil {
			log.Printf("Error sending chunk: %v", err)
		}

		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil{
		log.Printf("Upload failed: %v", err)
	}

	log.Printf("File uploaded, server message: %s", resp.Message)

}
