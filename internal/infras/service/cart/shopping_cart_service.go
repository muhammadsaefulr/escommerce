package service

import "github.com/muhammadsaefulr/escommerce/internal/domain/entity"

type ShoppingCartService struct {
	repo entity.ShoppingCartRepository
}

func NewShoppingCartService(repo entity.ShoppingCartRepository) *ShoppingCartService {
	return &ShoppingCartService{repo: repo}
}

func (s *ShoppingCartService) AddShoppingCart(cart *entity.ShoppingCart) (*entity.ShoppingCart, error) {
	return s.repo.AddShoppingCart(cart)
}

func (s *ShoppingCartService) GetShoppingCartById(ID string) (*entity.ShoppingCart, error) {
	return s.repo.GetShoppingCartById(ID)
}

func (s *ShoppingCartService) DeleteShoppingCart(ID string) error {
	return s.repo.DeleteShoppingCart(ID)
}
