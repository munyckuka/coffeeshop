package main

import (
	"api-gateway/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Настройка CORS
	router.Use(cors.Default())

	// Регистрация маршрутов
	routes.RegisterRoutes(router)

	// Запуск
	router.Run(":8088") // Gateway работает на порту 8088
}
