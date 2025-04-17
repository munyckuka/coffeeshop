package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"order-servive/domain"
	grpcHandler "order-servive/grpc"
	orderpb "order-servive/proto"
	"order-servive/repository"
	"order-servive/service"
)

func main() {
	// 1. Подключение к MongoDB
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("orders_db")

	// 2. Инициализация репозитория, сервиса и gRPC-хендлера
	var repo domain.OrderRepository = repository.NewMongoOrderRepository(db)
	orderService := service.NewOrderService(repo)
	orderHandler := grpcHandler.NewOrderGRPCHandler(orderService)

	// 3. Настройка gRPC-сервера
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("❌ Не удалось слушать порт: %v", err)
	}

	server := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(server, orderHandler)

	fmt.Println("✅ Order gRPC server запущен на :50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
