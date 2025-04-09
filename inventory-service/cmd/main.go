package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"inventory-service/delivery"
	"inventory-service/repository"
	"inventory-service/service"

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
	productRepo := repository.NewMongoProductRepository(db)
	productService := service.NewProductService(productRepo)

	// Настройка роутера
	router := gin.Default()
	delivery.NewProductHandler(router, productService)

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}
