package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSeller struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid" json:"id,omitempty"`
	UserID    uuid.UUID      `gorm:"unique;not null" json:"user_id" validate:"required"` // Menghubungkan penjual dengan pengguna
	NamaToko  string         `json:"nama_toko,omitempty" validate:"required,min=3,max=75"`
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
	CreateUserSeller(UserSeller *UserSeller) (*UserSeller, error)
	GetUserSellerById(id string) (*UserSeller, error)
	GetUserSellerByUserId(id string) (*UserSeller, error)
	GetUserByUserEmail(email string) (*User, error)
	UpdateUserSellerData(id string, UserSeller *UserSeller) error
	DeleteUserSellerById(id string) error
}
