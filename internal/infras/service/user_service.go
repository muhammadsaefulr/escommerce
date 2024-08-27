package service

import (
	domain "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	entity "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
)

type UserService struct {
	userRepo entity.UserRepository
}

func NewUserService(repo entity.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) AuthLoginUser(login *entity.AuthLoginUser) (*domain.User, error) {
	return s.userRepo.AuthLoginUser(login)
}

func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserById(id int) (*domain.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *UserService) UpdateUserData(id int, user *domain.User) error {
	return s.userRepo.UpdateUserData(id, user)
}

func (s *UserService) DeleteUserById(id int) error {
	return s.userRepo.DeleteUserById(id)
}
