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

func (s *UserSellerService) CreateUserSeller(UserSellerReg *domain.UserSellerRegister) (*domain.UserSeller, error) {
	return s.UserSellerRepo.CreateUserSeller(UserSellerReg)
}

func (s *UserSellerService) GetUserSellerById(id string) (*domain.UserSeller, error) {
	return s.UserSellerRepo.GetUserSellerById(id)
}

func (s *UserSellerService) GetUserSellerByUserId(id string) (*domain.UserSeller, error) {
	return s.UserSellerRepo.GetUserSellerByUserId(id)
}

func (s *UserSellerService) GetUserByUserEmail(email string) (*domain.User, error) {
	return s.UserSellerRepo.GetUserByUserEmail(email)
}

func (s *UserSellerService) UpdateUserSellerData(id string, UserSeller *domain.UserSeller) error {
	return s.UserSellerRepo.UpdateUserSellerData(id, UserSeller)
}

func (s *UserSellerService) DeleteUserSellerById(id string) error {
	return s.UserSellerRepo.DeleteUserSellerById(id)
}
