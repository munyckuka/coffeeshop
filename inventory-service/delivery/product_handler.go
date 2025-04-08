package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"inventory-service/domain"
	"inventory-service/service"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(router *gin.Engine, ps *service.ProductService) {
	handler := &ProductHandler{productService: ps}

	products := router.Group("/products")
	{
		products.POST("/", handler.CreateProduct)
		products.GET("/", handler.ListProducts)
		products.GET("/:id", handler.GetProduct)
		products.PATCH("/:id", handler.UpdateProduct)
		products.DELETE("/:id", handler.DeleteProduct)
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.CreateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.ID = id
	if err := h.productService.UpdateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.productService.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	// Пока фильтрация простая — пустой фильтр
	products, err := h.productService.ListProducts(map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
