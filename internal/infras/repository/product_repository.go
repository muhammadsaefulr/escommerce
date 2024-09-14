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

func (r *ProductRepositoryImpl) GetAllProduct() ([]entity.ProductItems, error) {
	var product []entity.ProductItems

	err := r.DB.Preload("Seller").Find(&product).Error

	return product, err
}

func (r *ProductRepositoryImpl) GetProductItems(ID string) (*entity.ProductItems, error) {
	var product entity.ProductItems

	err := r.DB.Where("ID = ?", ID).First(&product).Error

	return &product, err
}

func (r *ProductRepositoryImpl) GetSellerById(ID string) (*entity.UserSeller, error) {
	var seller entity.UserSeller

	err := r.DB.Where("ID = ?", ID).First(&seller).Error

	return &seller, err
}

func (r *ProductRepositoryImpl) UpdateProductItems(ID string, product *entity.ProductItems, updatedProduct *entity.ProductItems) (*entity.ProductItems, error) {

	err := r.DB.Model(&product).Where("ID = ?", ID).Updates(updatedProduct).Error

	return product, err
}

func (r *ProductRepositoryImpl) DeleteProductItems(ID string) error {

	log.Println("Deleting product with ID:", ID)
	return r.DB.Where("ID = ?", ID).Delete(&entity.ProductItems{}).Error

}
