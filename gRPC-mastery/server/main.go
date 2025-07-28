package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/dmehra2102/grpc-mastery/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	users map[string]*pb.User
	mu    sync.RWMutex
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

	log.Printf("Server listening on %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
