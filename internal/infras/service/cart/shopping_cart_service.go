package service

import "github.com/muhammadsaefulr/escommerce/internal/domain/entity"

type ShoppingCartService struct {
	repo entity.ShoppingCartRepository
}

func NewShoppingCartService(repo entity.ShoppingCartRepository) *ShoppingCartService {
	return &ShoppingCartService{repo: repo}
}

func (s *ShoppingCartService) GetShoppingCartById(ID string) (*entity.ShoppingCart, error) {
	return s.repo.GetShoppingCartById(ID)
}

func (s *ShoppingCartService) DeleteShoppingCart(ID string) error {
	return s.repo.DeleteShoppingCart(ID)
}

func (s *ShoppingCartService) AddShoppingCartItem(cartItem *entity.ShoppingCartItems) (*entity.ShoppingCartItems, error) {
	return s.repo.AddShoppingCartItem(cartItem)
}

func (s *ShoppingCartService) GetShoppingCartItemById(cartID string, productId string) (*entity.ShoppingCartItems, error) {
	return s.repo.GetShoppingCartItemById(cartID, productId)
}

func (s *ShoppingCartService) UpdateShoppingCartItem(cartID string, cartItem *entity.ShoppingCartItems) (*entity.ShoppingCartItems, error) {
	return s.repo.UpdateShoppingCartItem(cartID, cartItem)
}
