package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"order-servive/domain"
	"order-servive/service"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(router *gin.Engine, os *service.OrderService) {
	handler := &OrderHandler{orderService: os}

	orders := router.Group("/orders")
	{
		orders.POST("/", handler.CreateOrder)
		orders.GET("/:id", handler.GetOrderByID)
		orders.PATCH("/:id", handler.UpdateOrder)
	}

	router.GET("/users/:userId/orders", handler.ListOrdersByUser)
}

// CreateOrder создает новый заказ
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrderByID получает заказ по ID
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder обновляет статус заказа
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.UpdateOrder(id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}

// ListOrdersByUser возвращает список заказов пользователя
func (h *OrderHandler) ListOrdersByUser(c *gin.Context) {
	userID := c.Param("userId")

	orders, err := h.orderService.ListByUserId(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
