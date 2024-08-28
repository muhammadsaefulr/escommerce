package service

import (
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
)

type ProductService struct {
	productRepo entity.ProductRepository
}

func NewProductService(repo entity.ProductRepository) *ProductService {
	return &ProductService{productRepo: repo}
}

func (s *ProductService) AddProductItems(product *entity.ProductItems) (*entity.ProductItems, error) {
	return s.productRepo.AddProductItems(product)
}

func (s *ProductService) GetProductItems(ID string) (*entity.ProductItems, error) {
	return s.productRepo.GetProductItems(ID)
}
