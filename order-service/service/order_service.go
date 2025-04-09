package service

import (
	"order-servive/domain"
)

type OrderService struct {
	repo domain.OrderRepository
}

func NewOrderService(repo domain.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(p *domain.Order) error {
	return s.repo.Create(p)
}

func (s *OrderService) GetOrderByID(id string) (*domain.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) UpdateOrder(id string, status string) error {
	return s.repo.UpdateStatus(id, status)
}

func (s *OrderService) ListByUserId(userID string) ([]*domain.Order, error) {
	return s.repo.ListByUser(userID)
}
