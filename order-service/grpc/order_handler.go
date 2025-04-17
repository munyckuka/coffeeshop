package grpc

import (
	"context"
	"time"

	"order-servive/domain"
	orderpb "order-servive/proto"
	"order-servive/service"
)

type OrderGRPCHandler struct {
	orderpb.UnimplementedOrderServiceServer
	service *service.OrderService
}

func NewOrderGRPCHandler(s *service.OrderService) orderpb.OrderServiceServer {
	return &OrderGRPCHandler{service: s}
}

func (h *OrderGRPCHandler) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	order := &domain.Order{
		UserID:    req.GetUserId(),
		Items:     convertOrderItemsFromProto(req.GetItems()),
		Status:    "pending",
		CreatedAt: time.Now().Unix(),
	}

	if err := h.service.CreateOrder(order); err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{Order: convertOrderToProto(order)}, nil
}

func (h *OrderGRPCHandler) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	order, err := h.service.GetOrderByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &orderpb.GetOrderResponse{Order: convertOrderToProto(order)}, nil
}

func (h *OrderGRPCHandler) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.UpdateOrderStatusResponse, error) {
	if err := h.service.UpdateOrder(req.GetId(), req.GetStatus()); err != nil {
		return nil, err
	}

	updated, err := h.service.GetOrderByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &orderpb.UpdateOrderStatusResponse{Order: convertOrderToProto(updated)}, nil
}

func (h *OrderGRPCHandler) ListOrdersByUser(ctx context.Context, req *orderpb.ListOrdersByUserRequest) (*orderpb.ListOrdersByUserResponse, error) {
	orders, err := h.service.ListByUserId(req.GetUserId())
	if err != nil {
		return nil, err
	}

	var pbOrders []*orderpb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, convertOrderToProto(order))
	}

	return &orderpb.ListOrdersByUserResponse{Orders: pbOrders}, nil
}

// ---------- Конвертация моделей ----------

func convertOrderToProto(o *domain.Order) *orderpb.Order {
	var items []*orderpb.OrderItem
	for _, item := range o.Items {
		items = append(items, &orderpb.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		})
	}

	return &orderpb.Order{
		Id:        o.ID,
		UserId:    o.UserID,
		Items:     items,
		Status:    o.Status,
		CreatedAt: o.CreatedAt,
	}
}

func convertOrderItemsFromProto(items []*orderpb.OrderItem) []domain.OrderItem {
	var result []domain.OrderItem
	for _, item := range items {
		result = append(result, domain.OrderItem{
			ProductID: item.GetProductId(),
			Quantity:  int(item.GetQuantity()),
		})
	}
	return result
}
