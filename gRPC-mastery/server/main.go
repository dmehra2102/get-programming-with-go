package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	pbStream "github.com/dmehra2102/grpc-mastery/proto/stream"
	pb "github.com/dmehra2102/grpc-mastery/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type streamServiceServer struct {
	pbStream.UnimplementedStreamServiceServer
}

type userServiceServer struct {
	mu    sync.RWMutex
	users map[string]*pb.User
	pb.UnimplementedUserServiceServer
}

func (s *userServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received GetUser request: ID=%s", req.GetId())

	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[req.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "User with ID %s not found", req.GetId())
	}

	log.Printf("Retrieved user: ID=%s, Name=%s", user.GetId(), user.GetName())
	return &pb.GetUserResponse{User: user}, nil
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Received CreateUser request: Name=%s, Email=%s", req.GetName(), req.GetEmail())

	if req.GetName() == "" || req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name and Email cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	newUser := &pb.User{
		Id:    fmt.Sprintf("user-%d", time.Now().UnixNano()),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	s.users[newUser.Id] = newUser
	log.Printf("Created user: ID=%s, Name=%s", newUser.GetId(), newUser.GetName())

	return &pb.CreateUserResponse{User: newUser}, nil
}

func (s *streamServiceServer) GetMessages(req *pbStream.StreamRequest, stream pbStream.StreamService_GetMessagesServer) error {
	log.Printf("Recieved ServerStreaming request : %s", req.GetMessage())
	for i := 0; i < 5; i++ {
		responseMessage := fmt.Sprintf("Server streaming response %d for: %s", i+1, req.GetMessage())
		if err := stream.Send(&pbStream.StreamResponse{Result: responseMessage}); err != nil {
			log.Printf("Error sending stream message : %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	log.Printf("Finished sending server stream")
	return nil
}

func (s *streamServiceServer) SendMessages(stream pbStream.StreamService_SendMessagesServer) error {
	var receivedMessages []string
	log.Println("Server: Starting to receive client stream messages...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Server received all client stream messages: %v", receivedMessages)
			return stream.SendAndClose(&pbStream.StreamResponse{Result: "Server received all messages successfully!"})
		}
		if err != nil {
			log.Printf("Error receiving client stream: %v", err)
			return err
		}
		log.Printf("Received client stream request: %s", req.GetMessage())
		receivedMessages = append(receivedMessages, req.GetMessage())
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	userService := &userServiceServer{
		users: make(map[string]*pb.User, 0),
	}
	pb.RegisterUserServiceServer(s, userService)

	streamService := &streamServiceServer{}
	pbStream.RegisterStreamServiceServer(s, streamService)

	log.Printf("Server listening on %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
