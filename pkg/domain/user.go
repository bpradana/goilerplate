package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
}

type UserUsecase interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Create(u *User) (*User, error)
	Update(id int, u *User) (*User, error)
	Delete(id int) error
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Create(u *User) (*User, error)
	Update(id int, u *User) (*User, error)
	Delete(id int) error
}
