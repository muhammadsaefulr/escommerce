package repository

import (
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	domain "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) entity.UserRepository {
	return &UserRepositoryImpl{DB}
}

func (r *UserRepositoryImpl) CreateUser(user *domain.User) (*domain.User, error) {

	err := r.DB.Create(user).Error

	return user, err
}

func (r *UserRepositoryImpl) AuthLoginUser(loginEntity *entity.AuthLoginUser) (*domain.User, error) {
	var user domain.User

	err := r.DB.Where("email = ?", loginEntity.Email).First(&user).Error

	return &user, err
}

func (r *UserRepositoryImpl) GetUserById(id int) (*domain.User, error) {
	var user domain.User

	err := r.DB.Where("ID = ?", id).First(&user).Error

	return &user, err
}

func (r *UserRepositoryImpl) UpdateUserData(id int, user *domain.User) error {

	return r.DB.Where("id = ?", id).Updates(&user).Error
}

func (r *UserRepositoryImpl) DeleteUserById(id int) error {

	return r.DB.Delete(&domain.User{}, id).Error
}
