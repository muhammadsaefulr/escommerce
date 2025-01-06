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

func (s *UserService) GetUserById(id string) (*domain.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *UserService) UpdateUserData(id string, userUpdate *domain.UpdateUserData) error {
	return s.userRepo.UpdateUserData(id, userUpdate)
}

func (s *UserService) DeleteUserById(id string) error {
	return s.userRepo.DeleteUserById(id)
}
