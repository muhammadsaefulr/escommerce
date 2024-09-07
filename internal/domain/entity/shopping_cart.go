package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID                uuid.UUID           `gorm:"primary_key"`
	UserId            int                 `json:"user_id"`
	ShoppingCartItems []ShoppingCartItems `json:"shopping_cart_items",gorm:"foreignKey:CartId"`
}

func (base *ShoppingCart) BeforeCreate(db *gorm.DB) error {

	uuids, err := uuid.NewV6()
	if err != nil {
		return err
	}

	if base.ID == uuid.Nil {
		base.ID = uuids
	}

	return nil
}

// type Interface For Function

type ShoppingCartRepository interface {
	AddShoppingCart(cart *ShoppingCart) (*ShoppingCart, error)
	UpdateShoppingCart(ID string, cart *ShoppingCart) (*ShoppingCart, error)
	GetShoppingCartById(ID string) (*ShoppingCart, error)
	DeleteShoppingCart(ID string) error
}
