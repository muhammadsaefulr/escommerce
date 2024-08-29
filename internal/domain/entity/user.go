package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `json:"name" validate:"required,min=3,max=75"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email,min=1,max=100"`
	Password  string    `json:"password" validate:"required,min=3"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Type Struct For Promise Method

type AuthLoginUser struct {
	Email    string `json:"email" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=3"`
}

// Type Interface For Functions

type UserRepository interface {
	AuthLoginUser(login *AuthLoginUser) (*User, error)
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	UpdateUserData(id int, user *User) error
	DeleteUserById(id int) error
}
