package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "productinfo/client/ecommerce"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:50051"
	defaultTimeout = time.Second
)

func main() {
	// Establish gRPC connection
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}()

	// Initialize clients
	productClient := pb.NewProductInfoClient(conn)
	orderClient := pb.NewOrderManagementClient(conn)

	// Product operations
	if err := productOperations(context.Background(), productClient); err != nil {
		log.Fatalf("Product operations failed: %v", err)
	}

	// Order operations
	if err := orderOperations(context.Background(), orderClient); err != nil {
		log.Fatalf("Order operations failed: %v", err)
	}
}

func productOperations(ctx context.Context, client pb.ProductInfoClient) error {
	// Prepare product data
	product := &pb.Product{
		Name:        "Apple iPhone 11",
		Description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode.",
		Price:       1000.0,
	}

	// Add product
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	productID, err := client.AddProduct(ctx, product)
	if err != nil {
		return fmt.Errorf("could not add product: %w", err)
	}
	log.Printf("Product ID: %s added successfully", productID.Value)

	// Get product
	ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	retrievedProduct, err := client.GetProduct(ctx, &pb.ProductId{Value: productID.Value})
	if err != nil {
		return fmt.Errorf("could not get product: %w", err)
	}
	log.Printf("Product details: %v", retrievedProduct)

	return nil
}

func orderOperations(ctx context.Context, client pb.OrderManagementClient) error {
	// Get single order
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	order, err := client.GetOrder(ctx, &pb.OrderId{Value: "106"})
	if err != nil {
		return fmt.Errorf("failed to get order 106: %w", err)
	}
	log.Printf("Order details: %v", order)

	// Search orders
	ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	searchStream, err := client.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Google"})
	if err != nil {
		return fmt.Errorf("search orders failed: %w", err)
	}

	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error receiving search result: %w", err)
		}
		log.Printf("Search result: %v", searchOrder)
	}

	return nil
}