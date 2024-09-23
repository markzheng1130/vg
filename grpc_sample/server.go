package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"io/ioutil"
	"vg/grpc_sample/greetpb" // Import the generated protobuf code

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	data, err := ioutil.ReadFile("./sample.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	name := string(data)
	message := fmt.Sprintf("Hello, %s!", name)
	res := &greetpb.GreetResponse{
		Message: message,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
