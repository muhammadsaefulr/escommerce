package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSeller struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid" json:"id,omitempty"`
	Name      string         `json:"name,omitempty" validate:"required,min=3,max=75"`
	Email     string         `json:"email,omitempty" gorm:"unique" validate:"required,email,min=1,max=100"`
	Password  string         `json:"password,omitempty" validate:"required,min=3"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	Products  []ProductItems `gorm:"foreignKey:SellerId" json:"products,omitempty"`
}

func (base *UserSeller) BeforeCreate(db *gorm.DB) error {
	uuids, err := uuid.NewV6()
	if err != nil {
		return err
	}

	if base.ID == uuid.Nil {
		base.ID = uuids
	}

	return nil
}

// Type Struct For Promise Method

type AuthLoginUserSeller struct {
	Email    string `json:"email" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=3"`
}

// Type Interface For Functions

type UserSellerRepository interface {
	AuthLoginUserSeller(login *AuthLoginUserSeller) (*UserSeller, error)
	CreateUserSeller(UserSeller *UserSeller) (*UserSeller, error)
	GetUserSellerByEmail(email string) (*UserSeller, error)
	GetUserSellerById(id string) (*UserSeller, error)
	UpdateUserSellerData(id string, UserSeller *UserSeller) error
	DeleteUserSellerById(id string) error
}
