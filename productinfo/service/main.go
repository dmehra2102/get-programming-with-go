package main

import (
	"log"
	"net"

	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

func main() {
	lis,err := net.Listen("tcp",PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	serverInstance := &server{}
	pb.RegisterProductInfoServer(s,serverInstance)
	pb.RegisterOrderManagementServer(s, serverInstance)
	log.Printf("Starting gRPC listener on port : %s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}