package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserRegister struct {
	Name     string `json:"name" validate:"required string"`
	Email    string `json:"email" validate:"required email"`
	Phone    string `json:"phone" validate:"required string"`
	Password string `json:"password" validate:"required string"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required email"`
	Password string `json:"password" validate:"required string"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Type  string `json:"type"`
}

type UserDetail struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Type struct {
	ID   string `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	Name string `json:"name" gorm:"unique"`
}

func (t *Type) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return
}

type UserType struct {
	ID     string `json:"id" gorm:"primaryKey, type:uid, default:uuid_generate_v4()"`
	UserID string `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	TypeID string `json:"type_id"`
	Type   Type   `json:"type" gorm:"foreignKey:TypeID"`
}

func (u *UserType) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

// User Token has related to User
type UserToken struct {
	ID     string `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID string `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Type   string `json:"type"`
	Token  string `json:"token"`
}

func (u *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserBalance struct {
	ID      string  `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID  string  `json:"user_id"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Balance float64 `json:"balance" default:"0"`
}

func (u *UserBalance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserBalanceRequest struct {
	Balance float64 `json:"balance" validate:"required"`
}
