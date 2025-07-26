package main

import (
	"context"
	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetOrder(ctx context.Context, in *pb.OrderId)(*pb.Order,error){
	value,exists := s.orderMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil , status.Errorf(codes.NotFound, "Order does not exist with id : %s.", in.Value)
}