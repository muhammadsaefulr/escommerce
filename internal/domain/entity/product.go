package entity

import (
	"time"

	"github.com/muhammadsaefulr/escommerce/scripts"
	"gorm.io/gorm"
)

type ProductItems struct {
	ID                 string      `json:"id" gorm:"primaryKey;type:char(8)"`
	ProductName        string      `json:"name" validate:"required"`
	ProductDescription string      `json:"description" validate:"required"`
	ProductPrice       float32     `json:"price" validate:"required"`
	CreatedAt          time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	CategoryId         uint        `json:"category_id" validate:"required" gorm:"not null"`
	SellerId           string      `json:"seller_id" validate:"required" gorm:"type:text"`
	Seller             *UserSeller `json:"seller,omitempty" validate:"-", gorm:"foreignKey:SellerId;embedded"`
}

type AddProductItems struct {
	ProductName        string  `json:"name" validate:"required"`
	ProductDescription string  `json:"description" validate:"required"`
	ProductPrice       float32 `json:"price" validate:"required"`
	CategoryId         uint    `json:"category_id" validate:"required"`
}

type UpdateProductItems struct {
	ProductName        string  `json:"name" validate:"omitempty"`
	ProductDescription string  `json:"description" validate:"omitempty"`
	ProductPrice       float32 `json:"price" validate:"omitempty"`
	CategoryId         uint    `json:"category_id" validate:"omitempty"`
}

type FilteredProductReturn struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	Price     float32    `json:"price"`
	SellerID  string     `json:"seller_id"`
}

func (base *ProductItems) BeforeCreate(db *gorm.DB) (err error) {

	id, err := scripts.GenerateUniqueHexId4B()

	if err != nil {
		return err
	}

	base.ID = id

	return

}

// Type Interface For Function

type ProductRepository interface {
	AddProductItems(product *ProductItems) (*ProductItems, error)
	GetAllProductBySellerId(sellerId string) ([]ProductItems, error)
	GetProductItems(ID string) (*ProductItems, error)
	GetSellerById(ID string) (*UserSeller, error)
	UpdateProductItems(ID string, product *ProductItems, updatedProduct *ProductItems) (*ProductItems, error)
	DeleteProductItems(ID string) error
}
