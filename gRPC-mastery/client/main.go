package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dmehra2102/grpc-mastery/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Uniary RPC : CreateUser
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createUserReq := &pb.CreateUserRequest{
		Name:  "Deepanshu Mehra",
		Email: "deepanshumehra2102@gmail.com",
	}

	createUserResp, err := client.CreateUser(ctx, createUserReq)
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}
	log.Printf("Created User: ID=%s, Name=%s, Email=%s",
		createUserResp.GetUser().GetId(),
		createUserResp.GetUser().GetName(),
		createUserResp.GetUser().GetEmail())

	// Unary RPC : GetUser
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getUserReq := &pb.GetUserRequest{
		Id: createUserResp.GetUser().GetId(),
	}
	getUserResp, err := client.GetUser(ctx, getUserReq)
	if err != nil {
		log.Fatalf("Could not get user: %v", err)
	}
	log.Printf("Retrieved User: ID=%s, Name=%s, Email=%s",
		getUserResp.GetUser().GetId(),
		getUserResp.GetUser().GetName(),
		getUserResp.GetUser().GetEmail())

	// --- Unary RPC: GetUser (Not Found Example) ---
	ctx,cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getUserNotFoundReq := &pb.GetUserRequest{
		Id: "non-existent-id",
	}
	_, err = client.GetUser(ctx, getUserNotFoundReq)
	if err!= nil {
		log.Printf("Attempted to get non-existent user, got expected error: %v", err)
	}
}
