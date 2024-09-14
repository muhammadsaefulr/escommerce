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

func (s *ProductService) GetAllProduct() ([]entity.ProductItems, error) {
	return s.productRepo.GetAllProduct()
}

func (s *ProductService) GetProductItems(ID string) (*entity.ProductItems, error) {
	return s.productRepo.GetProductItems(ID)
}

func (s *ProductService) GetSellerById(ID string) (*entity.UserSeller, error) {
	return s.productRepo.GetSellerById(ID)
}

func (s *ProductService) UpdateProductItems(ID string, product *entity.ProductItems, updatedProduct *entity.ProductItems) (*entity.ProductItems, error) {
	return s.productRepo.UpdateProductItems(ID, product, updatedProduct)
}

func (s *ProductService) DeleteProductItems(ID string) error {
	return s.productRepo.DeleteProductItems(ID)
}
