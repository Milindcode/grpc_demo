package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	sample "github.com/Milindcode/grpc_server_streaming/proto"
	"google.golang.org/grpc"
)

type StockServer struct {
	sample.UnimplementedStockServiceServer
}

func (s *StockServer) TestFunc(req *sample.StockRequest, stream sample.StockService_TestFuncServer) error {
	log.Println("Recieving Stock Request for symbol: " + req.Symbol)

	for i := 1; i <= 10; i++ {
		log.Printf("Sending tick data: %v\n", i)

		if err := stream.Send(&sample.StockResponse{
			TimeStamp: int64(i),
			Price:     rand.Int63n(10000),
			Symbol:    req.Symbol,
		}); err != nil {
			log.Printf("Server failed to deliver data: %v\n", err)
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("Cant listen: %v\n", err)
	}

	grpc_server := grpc.NewServer()
	sample.RegisterStockServiceServer(grpc_server, &StockServer{})

	log.Println("Server starting at port :50051")
	err = grpc_server.Serve(listener)
	if err != nil {
		log.Printf("Error serving : %v\n", err)
	}
}
