package repository

import (
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/gorm"
)

type UserSellerRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserSellerRepository(DB *gorm.DB) entity.UserSellerRepository {
	return &UserSellerRepositoryImpl{DB}
}

func (r *UserSellerRepositoryImpl) CreateUserSeller(UserSellerRegist *entity.UserSellerRegister) (*entity.UserSeller, error) {
	var userSeller entity.UserSeller

	userSeller.UserID = UserSellerRegist.UserID
	userSeller.NamaToko = UserSellerRegist.NamaToko

	err := r.DB.Create(&userSeller).Error

	return &userSeller, err
}

func (r *UserSellerRepositoryImpl) GetUserByUserEmail(email string) (*entity.User, error) {
	var User entity.User

	err := r.DB.First(&User, "email = ?", email).Error

	return &User, err
}

func (r *UserSellerRepositoryImpl) GetUserSellerByUserId(id string) (*entity.UserSeller, error) {
	var UserSeller entity.UserSeller

	err := r.DB.First(&UserSeller, "user_id = ?", id).Error

	return &UserSeller, err
}

func (r *UserSellerRepositoryImpl) GetUserSellerById(id string) (*entity.UserSeller, error) {
	var UserSeller entity.UserSeller

	err := r.DB.Preload("Products").First(&UserSeller, "id = ?", id).Error

	return &UserSeller, err
}

func (r *UserSellerRepositoryImpl) UpdateUserSellerData(id string, UserSeller *entity.UserSeller) error {

	return r.DB.Where("id = ?", id).Updates(&UserSeller).Error
}

func (r *UserSellerRepositoryImpl) DeleteUserSellerById(id string) error {

	return r.DB.Delete(&entity.UserSeller{}, "id = ?", id).Error
}
