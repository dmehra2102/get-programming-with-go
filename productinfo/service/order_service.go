package main

import (
	"context"
	"fmt"
	"log"
	pb "productinfo/service/ecommerce"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *server) GetOrder(ctx context.Context, in *pb.OrderId) (*pb.Order, error) {
	value, exists := s.orderMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Order does not exist with id : %s.", in.Value)
}

func (s *server) SearchOrders(searchQuery *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)
		for _, item := range order.Items {
			if strings.Contains(item, searchQuery.Value) {
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
			}
		}
	}
	return nil
}
