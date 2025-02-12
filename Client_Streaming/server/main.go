package main

import (
	"io"
	"log"
	"net"
	"os"

	sample "github.com/Milindcode/grpc_client_streaming/proto"
	"google.golang.org/grpc"
)

type FileServer struct {
	sample.UnimplementedFileStreamerServer
}

func (server *FileServer) SendFile(stream sample.FileStreamer_SendFileServer) error {
	var file *os.File
	var fileName string

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			if file != nil {
				file.Close()
			}

			return stream.SendAndClose(&sample.ServerResponse{
				Status:  true,
				Message: "File recieved successfully",
			})
		}

		if err != nil {
			return err
		}

		if file == nil {
			fileName = chunk.FileName
			file, err = os.Create(fileName)
			if err != nil {
				return err
			}

			log.Println("Recieveing file: " + fileName)
		}

		_, err = file.Write(chunk.Chunk)
		if err != nil {
			return err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	sample.RegisterFileStreamerServer(grpcServer, &FileServer{})

	log.Println("Server starting at port: 50051")
	err = grpcServer.Serve(listener); if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
