package grpc

import (
	"context"
	inventorypb "inventory-service/proto"

	"github.com/google/uuid"
	"inventory-service/domain"
	"inventory-service/service"
)

type InventoryGrpcHandler struct {
	inventorypb.UnimplementedInventoryServiceServer
	productService *service.ProductService
}

func NewInventoryGrpcHandler(ps *service.ProductService) *InventoryGrpcHandler {
	return &InventoryGrpcHandler{
		productService: ps,
	}
}

func (h *InventoryGrpcHandler) CreateProduct(ctx context.Context, req *inventorypb.CreateProductRequest) (*inventorypb.CreateProductResponse, error) {
	p := req.GetProduct()
	product := &domain.Product{
		ID:       uuid.New().String(),
		Name:     p.GetName(),
		Category: p.GetCategoryId(),
		Price:    p.GetPrice(),
		Stock:    int(p.GetQuantity()),
	}

	err := h.productService.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &inventorypb.CreateProductResponse{
		Product: toProtoProduct(product),
	}, nil
}

func (h *InventoryGrpcHandler) GetProduct(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.GetProductResponse, error) {
	product, err := h.productService.GetProductByID(req.GetId())
	if err != nil {
		return nil, err
	}
	return &inventorypb.GetProductResponse{
		Product: toProtoProduct(product),
	}, nil
}

func (h *InventoryGrpcHandler) UpdateProduct(ctx context.Context, req *inventorypb.UpdateProductRequest) (*inventorypb.UpdateProductResponse, error) {
	p := req.GetProduct()
	product := &domain.Product{
		ID:       p.GetId(),
		Name:     p.GetName(),
		Category: p.GetCategoryId(),
		Price:    p.GetPrice(),
		Stock:    int(p.GetQuantity()),
	}

	err := h.productService.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return &inventorypb.UpdateProductResponse{
		Product: toProtoProduct(product),
	}, nil
}

func (h *InventoryGrpcHandler) DeleteProduct(ctx context.Context, req *inventorypb.DeleteProductRequest) (*inventorypb.DeleteProductResponse, error) {
	err := h.productService.DeleteProduct(req.GetId())
	if err != nil {
		return nil, err
	}
	return &inventorypb.DeleteProductResponse{Message: "Product deleted"}, nil
}

func (h *InventoryGrpcHandler) ListProducts(ctx context.Context, req *inventorypb.ListProductsRequest) (*inventorypb.ListProductsResponse, error) {
	products, err := h.productService.ListProducts(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var protoProducts []*inventorypb.Product
	for _, p := range products {
		protoProducts = append(protoProducts, toProtoProduct(p))
	}

	return &inventorypb.ListProductsResponse{
		Products: protoProducts,
	}, nil
}

func toProtoProduct(p *domain.Product) *inventorypb.Product {
	return &inventorypb.Product{
		Id:         p.ID,
		Name:       p.Name,
		CategoryId: p.Category,
		Price:      p.Price,
		Quantity:   int32(p.Stock),
	}
}
