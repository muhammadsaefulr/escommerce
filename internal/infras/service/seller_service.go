package service

import (
	domain "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	entity "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
)

type UserSellerService struct {
	UserSellerRepo entity.UserSellerRepository
}

func NewUserSellerService(repo entity.UserSellerRepository) *UserSellerService {
	return &UserSellerService{UserSellerRepo: repo}
}

func (s *UserSellerService) AuthLoginUserSeller(login *entity.AuthLoginUserSeller) (*domain.UserSeller, error) {
	return s.UserSellerRepo.AuthLoginUserSeller(login)
}

func (s *UserSellerService) CreateUserSeller(UserSeller *domain.UserSeller) (*domain.UserSeller, error) {
	return s.UserSellerRepo.CreateUserSeller(UserSeller)
}

func (s *UserSellerService) GetUserSellerById(id string) (*domain.UserSeller, error) {
	return s.UserSellerRepo.GetUserSellerById(id)
}

func (s *UserSellerService) GetUserSellerByEmail(email string) (*domain.UserSeller, error) {
	return s.UserSellerRepo.GetUserSellerByEmail(email)
}

func (s *UserSellerService) UpdateUserSellerData(id string, UserSeller *domain.UserSeller) error {
	return s.UserSellerRepo.UpdateUserSellerData(id, UserSeller)
}

func (s *UserSellerService) DeleteUserSellerById(id string) error {
	return s.UserSellerRepo.DeleteUserSellerById(id)
}
