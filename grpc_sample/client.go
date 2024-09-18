package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"vg/grpc_sample/greetpb" // Import the generated protobuf code

	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	// Call the Greet method
	req := &greetpb.GreetRequest{
		Name: "John Doe",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	GRPCMaxCallRecvMsgSizeInBytes := 1000000
	res, err := client.Greet(ctx, req, grpc.MaxCallRecvMsgSize(GRPCMaxCallRecvMsgSizeInBytes))
	if err != nil {
		log.Fatalf("Error calling Greet: %v", err)
	}

	fmt.Println("Response from server:", res.GetMessage())
}
