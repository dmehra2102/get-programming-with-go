package main

import (
	"context"
	pb "grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedStockServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error){
	log.Printf("Received : %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()},nil
}

func main() {
	lis,err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}
	s := grpc.NewServer()
	serverInstance := &server{}
	pb.RegisterGreeterServer(s,serverInstance)
	pb.RegisterStockServiceServer(s,serverInstance)
	log.Println("Starting gRPC listener on port : 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}