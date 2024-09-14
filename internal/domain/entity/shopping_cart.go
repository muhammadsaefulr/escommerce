package entity

import (
	"log"

	"github.com/muhammadsaefulr/escommerce/scripts"
	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID                string              `json:"id" gorm:"primaryKey"`
	UserId            string              `json:"user_id" validate:"required" gorm:"type:uuid"`
	ShoppingCartItems []ShoppingCartItems `json:"shopping_cart_items" gorm:"foreignKey:CartId"`
}

func (s *ShoppingCart) BeforeCreate(tx *gorm.DB) (err error) {
	generatedHex, err := scripts.GenerateUniqueHexId3B()

	if err != nil {
		log.Fatalf(err.Error())
		return err
	}

	if s.ID == "" {
		s.ID = generatedHex
	}

	return nil
}

// type Interface For Function

type ShoppingCartRepository interface {
	UpdateShoppingCart(ID string, cart *ShoppingCart) (*ShoppingCart, error)
	GetShoppingCartById(ID string) (*ShoppingCart, error)
	DeleteShoppingCart(ID string) error

	GetShoppingCartItemById(cartID string, productId string) (*ShoppingCartItems, error)
	AddShoppingCartItem(cartItem *ShoppingCartItems) (*ShoppingCartItems, error)
	UpdateShoppingCartItem(cartID string, cartItem *ShoppingCartItems) (*ShoppingCartItems, error)
	DeleteShoppingCartItem(cartID string) error
}
