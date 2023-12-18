package main

import (
	"context"
	"log"

	sopropb "github.com/sopro-dev/sopro-proto/pkg/sopro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := sopropb.NewSoproServiceClient(conn)

	// Example: Call the PingPong RPC
	healthReq := &sopropb.HealthRequest{}
	healthResp, err := client.PingPong(context.Background(), healthReq)
	if err != nil {
		log.Fatalf("Error calling PingPong: %v", err)
	}
	log.Printf("PingPong Response: %s", healthResp.Message)
}
