package main

import (
	"context"
	"log"
	"net"

	"github.com/Milindcode/grpc_demo/proto/samplepb"
	"google.golang.org/grpc"
)

type GreeterServer struct{
	samplepb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, in *samplepb.HelloRequest) (*samplepb.HelloResponse, error) {
	return &samplepb.HelloResponse{
		Greeting: "Hello, " + in.GetName(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	samplepb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	log.Println("Server starting at port: 50051")
	err = grpcServer.Serve(listener); if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
