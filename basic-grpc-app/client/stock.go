package main

import (
	"context"
	pb "grpc-app/proto"
	"io"
	"log"
)

func stockClientOperation(client pb.StockServiceClient) {
	req := &pb.StockRequest{Symbol: "GMR INFRA"}

	stream,err := client.GetStockPrice(context.Background(),req)
	if err!= nil {
		log.Fatalf("could not get stock prices: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server finished streaming.")
			break
		}
		if err!= nil {
			log.Fatalf("error receiving message: %v", err)
		}
		log.Printf("Received update: %s - $%.2f at %s", res.GetSymbol(), res.GetPrice(), res.GetTimestamp()) 
	}
	log.Println("Client finished receiving stock prices.")
}