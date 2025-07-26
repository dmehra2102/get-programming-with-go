package main

import (
	"context"
	pb "productinfo/service/ecommerce"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*pb.Product
	orderMap   map[string]*pb.Order
	pb.UnimplementedProductInfoServer
	pb.UnimplementedOrderManagementServer
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductId, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID : %v", err)
	}
	in.Id = id.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in
	return &pb.ProductId{Value: in.Id}, status.New(codes.OK, "Product Added").Err()
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *pb.ProductId) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist with id : %s.", in.Value)
}
