package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"order-servive/delivery"
	"order-servive/repository"
	"order-servive/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Подключение к MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	db := client.Database("ecommerce")

	// Инициализация слоёв
	orderRepo := repository.NewMongoOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)

	// Настройка роутера
	router := gin.Default()
	delivery.NewOrderHandler(router, orderService)

	// Запуск сервера
	if err := router.Run(":8081"); err != nil { // Порт можно поменять, если нужно
		log.Fatal("Server error:", err)
	}
}
