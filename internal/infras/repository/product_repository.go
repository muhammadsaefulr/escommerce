package repository

import (
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) entity.ProductRepository {
	return &ProductRepositoryImpl{DB}
}

func (r *ProductRepositoryImpl) AddProductItems(product *entity.ProductItems) (*entity.ProductItems, error) {
	err := r.DB.Create(product).Error

	return product, err
}

func (r *ProductRepositoryImpl) GetProductItems(ID string) (*entity.ProductItems, error) {
	var product entity.ProductItems

	err := r.DB.Where("ID = ?", ID).First(&product).Error

	return &product, err
}
