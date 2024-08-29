package entity

import (
	"gorm.io/gorm"
)

type CategoryProduct struct {
	gorm.Model
	Name        string         `json:"name" validate:"required,min=3,max=75"`
	Description string         `json:"description" validate:"required,min=3,max=75"`
	Products    []ProductItems `gorm:"foreignKey:CategoryId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Type Struct For Promise Method

// Type Interface For Functions

type CategoryRepository interface {
	CreateCategory(Category *CategoryProduct) (*CategoryProduct, error)
	GetCategoryById(id int) (*CategoryProduct, error)
	GetCategoryByName(categoryname string) (*CategoryProduct, error)
	UpdateCategoryData(id int, Category *CategoryProduct) error
	DeleteCategoryById(id int) error
}
