package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Name      string    `json:"name" validate:"required,min=3,max=75"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email,min=1,max=100"`
	Password  string    `json:"password" validate:"required,min=3"`
	RoleId    int       `json:"role_id" gorm:"default:2"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (base *User) BeforeCreate(db *gorm.DB) error {
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

type AuthLoginUser struct {
	Email    string `json:"email" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=3"`
}

type AuthRegisterUser struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=3"`
}

type UserDataReturnViews struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleId int    `json:"role_id"`
}

// Type Interface For Functions

type UserRepository interface {
	AuthLoginUser(login *AuthLoginUser) (*User, error)
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	UpdateUserData(id string, user *User) error
	DeleteUserById(id string) error
}
