package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/muhammadsaefulr/escommerce/scripts"
	"gorm.io/gorm"
)

type UserSeller struct {
	ID        string          `gorm:"primaryKey;type:text" json:"seller_id,omitempty"`
	UserID    uuid.UUID       `gorm:"unique;not null;type: uuid" json:"user_id" validate:"required"` // Menghubungkan penjual dengan pengguna
	NamaToko  string          `json:"nama_toko,omitempty" validate:"required,min=3,max=75"`
	CreatedAt time.Time       `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	Products  *[]ProductItems `gorm:"foreignKey:ID" json:"products,omitempty"`
}

type UserSellerRegister struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	NamaToko string    `json:"nama_toko,omitempty" validate:"required,min=3,max=75"`
}

func (base *UserSeller) BeforeCreate(db *gorm.DB) error {
	uniqueHexId, err := scripts.GenerateUniqueHexId6B()
	if err != nil {
		return err
	}

	if base.ID == "" {
		base.ID = uniqueHexId
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
	CreateUserSeller(UserSeller *UserSellerRegister) (*UserSeller, error)
	GetUserSellerById(id string) (*UserSeller, error)
	GetUserSellerByUserId(id string) (*UserSeller, error)
	GetUserByUserEmail(email string) (*User, error)
	UpdateUserSellerData(id string, UserSeller *UserSeller) error
	DeleteUserSellerById(id string) error
}
