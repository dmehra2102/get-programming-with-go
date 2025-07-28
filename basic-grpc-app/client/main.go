package main

import (
	"context"
	"log"
	"time"

	pb "grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	stockClient := pb.NewStockServiceClient(conn)


	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp,err := c.SayHello(ctx, &pb.HelloRequest{Name: "Deepanshu Mehra"})
	if err!= nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())

	// Stock Operation
	stockClientOperation(stockClient)
}