package repository

import (
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/gorm"
)

type ShoppingCartRepositoryImpl struct {
	DB *gorm.DB
}

func NewShoppingCartRepository(db *gorm.DB) entity.ShoppingCartRepository {
	return &ShoppingCartRepositoryImpl{DB: db}
}

func (r *ShoppingCartRepositoryImpl) UpdateShoppingCart(ID string, cart *entity.ShoppingCart) (*entity.ShoppingCart, error) {
	return nil, nil
}

func (r *ShoppingCartRepositoryImpl) GetShoppingCartById(cartID string) (*entity.ShoppingCart, error) {
	var cart entity.ShoppingCart

	err := r.DB.Where("user_id = ?", cartID).Preload("ShoppingCartItems.Product.Seller").First(&cart).Error

	return &cart, err
}

func (r *ShoppingCartRepositoryImpl) DeleteShoppingCart(ID string) error {
	return nil
}

//

func (r *ShoppingCartRepositoryImpl) GetShoppingCartItemById(cartID string, productId string) (*entity.ShoppingCartItems, error) {
	var cartItem entity.ShoppingCartItems

	err := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productId).First(&cartItem).Error

	return &cartItem, err
}

func (r *ShoppingCartRepositoryImpl) AddShoppingCartItem(cartItem *entity.ShoppingCartItems) (*entity.ShoppingCartItems, error) {
	err := r.DB.Model(entity.ShoppingCartItems{}).Create(&cartItem).Error

	return cartItem, err
}

func (r *ShoppingCartRepositoryImpl) UpdateShoppingCartItem(cartID string, cartItem *entity.ShoppingCartItems) (*entity.ShoppingCartItems, error) {
	err := r.DB.Save(cartItem).Error

	return cartItem, err
}

func (r *ShoppingCartRepositoryImpl) DeleteShoppingCartItem(cartID string) error {
	err := r.DB.Where("ID = ?", cartID).Delete(&entity.ShoppingCartItems{}).Error

	return err
}
