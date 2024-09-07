package entity

import "github.com/google/uuid"

type ShoppingCartItems struct {
	ID              int          `gorm:"primary_key"`
	CartId          uuid.UUID    `json:"cart_id" validate:"required,min=1"`
	ShoppingCart    ShoppingCart `gorm:"foreignKey;CartId;constrait:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductId       uuid.UUID    `validate:"required"`
	Product         ProductItems `gorm:"foreignKey,ProductId;constrait:CASCADE,OnDelete:SET NULL"`
	Quantity        int
	PriceAtPurchase int
}
