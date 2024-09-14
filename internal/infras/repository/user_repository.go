package repository

import (
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) entity.UserRepository {
	return &UserRepositoryImpl{DB}
}

func (r *UserRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	cart := entity.ShoppingCart{
		UserId: user.ID.String(),
	}

	if err := r.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) AuthLoginUser(loginEntity *entity.AuthLoginUser) (*entity.User, error) {
	var user entity.User

	err := r.DB.Where("email = ?", loginEntity.Email).First(&user).Error

	return &user, err
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *UserRepositoryImpl) GetUserById(id string) (*entity.User, error) {
	var user entity.User

	err := r.DB.Where("ID = ?", id).First(&user).Error

	return &user, err
}

func (r *UserRepositoryImpl) UpdateUserData(id string, user *entity.User) error {

	return r.DB.Where("id = ?", id).Updates(&user).Error
}

func (r *UserRepositoryImpl) DeleteUserById(id string) error {

	return r.DB.Delete(&entity.User{}, id).Error
}
