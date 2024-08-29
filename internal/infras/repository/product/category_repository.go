package repository

import (
	"log"

	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) entity.CategoryRepository {
	return &CategoryRepositoryImpl{DB}
}

func (r *CategoryRepositoryImpl) CreateCategory(Category *entity.CategoryProduct) (*entity.CategoryProduct, error) {

	err := r.DB.Create(Category).Error

	return Category, err
}

func (r *CategoryRepositoryImpl) GetCategoryById(id int) (*entity.CategoryProduct, error) {
	var Category entity.CategoryProduct

	log.Println(id)

	err := r.DB.Preload("Products").Where("id = ?", id).First(&Category).Error

	return &Category, err
}

func (r *CategoryRepositoryImpl) GetCategoryByName(categoryname string) (*entity.CategoryProduct, error) {
	var Category entity.CategoryProduct

	err := r.DB.Where("name = ?", categoryname).First(&Category).Error

	return &Category, err
}

func (r *CategoryRepositoryImpl) UpdateCategoryData(id int, Category *entity.CategoryProduct) error {

	log.Println(id)
	return r.DB.Where("ID = ?", id).Updates(&Category).Error
}

func (r *CategoryRepositoryImpl) DeleteCategoryById(id int) error {

	return r.DB.Delete(&entity.CategoryProduct{}, id).Error
}
