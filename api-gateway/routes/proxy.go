package routes

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	inventoryURL, _ := url.Parse("http://inventory-service:8080")
	orderURL, _ := url.Parse("http://order-service:8081")

	inventoryProxy := httputil.NewSingleHostReverseProxy(inventoryURL)
	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)

	router.Any("/products/*proxyPath", gin.WrapH(inventoryProxy))
	router.Any("/orders/*proxyPath", gin.WrapH(orderProxy))
	router.Any("/users/*proxyPath", gin.WrapH(orderProxy))
}
