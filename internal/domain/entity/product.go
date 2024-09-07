package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductItems struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ProductName        string    `json:"name" validate:"required"`
	ProductDescription string    `json:"description" validate:"required"`
	ProductPrice       float32   `json:"price" validate:"required"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CategoryId         uint      `json:"category_id" gorm:"not null"`
}

func (base *ProductItems) BeforeCreate(db *gorm.DB) error {

	uuids, err := uuid.NewV6()
	if err != nil {
		return err
	}

	if base.ID == uuid.Nil {
		base.ID = uuids
	}

	return nil
}

// Type Interface For Function

type ProductRepository interface {
	AddProductItems(product *ProductItems) (*ProductItems, error)
	GetProductItems(ID string) (*ProductItems, error)
	UpdateProductItems(ID string, product *ProductItems, updatedProduct *ProductItems) (*ProductItems, error)
	DeleteProductItems(ID string) error
}
