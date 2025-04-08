package service

import (
	"inventory-service/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(p *domain.Product) error {
	return s.repo.Create(p)
}

func (s *ProductService) GetProductByID(id string) (*domain.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) UpdateProduct(p *domain.Product) error {
	return s.repo.Update(p)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *ProductService) ListProducts(filter map[string]interface{}) ([]*domain.Product, error) {
	return s.repo.List(filter)
}
