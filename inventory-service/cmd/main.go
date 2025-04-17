package main

import (
	"context"
	inventorypb "inventory-service/proto"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	grpcHandler "inventory-service/grpc"

	"inventory-service/repository"
	"inventory-service/service"
)

func main() {
	// Подключение к MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	db := client.Database("inventorydb")
	repo := repository.NewMongoProductRepository(db)
	productService := service.NewProductService(repo)
	handler := grpcHandler.NewInventoryGrpcHandler(productService)

	// Запуск gRPC-сервера
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(grpcServer, handler)

	log.Println("✅ Inventory gRPC server is running on port 8080...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
