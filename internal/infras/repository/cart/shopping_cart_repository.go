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

func (r *ShoppingCartRepositoryImpl) AddShoppingCart(cart *entity.ShoppingCart) (*entity.ShoppingCart, error) {
	return nil, nil
}

func (r *ShoppingCartRepositoryImpl) UpdateShoppingCart(ID string, cart *entity.ShoppingCart) (*entity.ShoppingCart, error) {
	return nil, nil
}

func (r *ShoppingCartRepositoryImpl) GetShoppingCartById(ID string) (*entity.ShoppingCart, error) {
	return nil, nil
}

func (r *ShoppingCartRepositoryImpl) DeleteShoppingCart(ID string) error {
	return nil
}
