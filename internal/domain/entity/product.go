package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductItems struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	ProductName        string    `json:"name"`
	ProductDescription string    `json:"description"`
	ProductPrice       float32   `json:"price"`
	Created_at         time.Time `json:"created_at"`
	Updated_at         time.Time `json:"updated_at"`
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
	// UpdateProductItems(ID string, product *ProductItems) (*ProductItems, error)
	// DeleteProductItems(ID string) error
}
