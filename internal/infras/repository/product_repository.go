package repository

import (
	"log"

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

func (r *ProductRepositoryImpl) UpdateProductItems(ID string, product *entity.ProductItems) (*entity.ProductItems, error) {

	err := r.DB.Model(&entity.ProductItems{}).Where("ID = ?", ID).Save(&product).Error

	return product, err
}

func (r *ProductRepositoryImpl) DeleteProductItems(ID string) error {

	log.Println("Deleting product with ID:", ID)
	return r.DB.Where("ID = ?", ID).Delete(&entity.ProductItems{}).Error

}
