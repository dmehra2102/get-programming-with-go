package main

import (
	pb "grpc-app/proto"
	"log"
	"math/rand"
	"time"
)

func (s *server) GetStockPrice(req *pb.StockRequest, stream pb.StockService_GetStockPriceServer) error {
	log.Printf("Recieved request for stock prices for symbol : %v", req.GetSymbol())
	for i := 0 ; i<5;i++ {
		price := 100.0 + rand.Float64()*10.0
		timestamp := time.Now().Format(time.RFC3339)

		resp := &pb.StockResponse{
			Symbol: req.GetSymbol(),
			Price: price,
			Timestamp: timestamp,
		}

		if err := stream.Send(resp); err != nil {
			log.Printf("Error sending stock price update %d: %v", i+1, err)
			return err
		}
		log.Printf("Sent: %s - $%.2f at %s", resp.Symbol, resp.Price, resp.Timestamp)
		time.Sleep(time.Second)
	}
	log.Printf("Finished streaming stock prices for symbol: '%v'", req.GetSymbol())
	return nil
}