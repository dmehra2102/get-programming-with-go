package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pbStream "github.com/dmehra2102/grpc-mastery/proto/stream"
	pb "github.com/dmehra2102/grpc-mastery/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)
	streamClient := pbStream.NewStreamServiceClient(conn)

	// Uniary RPC : CreateUser
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createUserReq := &pb.CreateUserRequest{
		Name:  "Deepanshu Mehra",
		Email: "deepanshumehra2102@gmail.com",
	}

	createUserResp, err := userClient.CreateUser(ctx, createUserReq)
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
	getUserResp, err := userClient.GetUser(ctx, getUserReq)
	if err != nil {
		log.Fatalf("Could not get user: %v", err)
	}
	log.Printf("Retrieved User: ID=%s, Name=%s, Email=%s",
		getUserResp.GetUser().GetId(),
		getUserResp.GetUser().GetName(),
		getUserResp.GetUser().GetEmail())

	// --- Unary RPC: GetUser (Not Found Example) ---
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getUserNotFoundReq := &pb.GetUserRequest{
		Id: "non-existent-id",
	}
	_, err = userClient.GetUser(ctx, getUserNotFoundReq)
	if err != nil {
		log.Printf("Attempted to get non-existent user, got expected error: %v", err.Error())
	}

	// Server-Streaming RPC
	log.Println("\n--- Server Streaming RPC: GetMessages ---")
	req := &pbStream.StreamRequest{Message: "Hello from client for server stream"}
	stream, err := streamClient.GetMessages(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not get messages: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server streaming finished on client side.")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		log.Printf("Recieved from server : %s", resp.GetResult())
	}

	// Client-Streaming RPC
	log.Println("\n--- Client Streaming RPC: SendMessages ---")
	clientStream, err := streamClient.SendMessages(context.Background())
	if err != nil {
		log.Fatalf("Could not open client stream: %v", err)
	}

	messagesToSend := []string{"First", "Second", "Third", "Fourth", "Fifth"}
	for i, msg := range messagesToSend {
		req := &pbStream.StreamRequest{Message: fmt.Sprintf("Client Stream Request %d: %s", i+1, msg)}
		if err := clientStream.Send(req); err != nil {
			log.Fatalf("Error sending client stream message: %v", err)
		}
		log.Printf("Client sent: %s", req.GetMessage())
		time.Sleep(200 * time.Millisecond)
	}

	res,err := clientStream.CloseAndRecv()
	if err!= nil {
		log.Fatalf("Error closing client stream and receiving response: %v", err)
	}
	log.Printf("Client Streaming Response: %s", res.GetResult())
}
