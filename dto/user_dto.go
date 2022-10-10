package dto

import (
	"dataekspor-be/entity"
	"time"
)

type UserUpdateDTO struct {
	ID        string               `json:"id"`
	Name      string               `json:"name"`
	Email     string               `json:"email"`
	Phone     string               `json:"telepon"`
	Password  string               `json:"password"`
	Role      entity.SliceOfString `json:"role"`
	UpdatedAt time.Time           `json:"updatedAt"`
}

type UserRegisterDTO struct {
	Name      string               `json:"name" binding:"required"`
	Email     string               `json:"email" binding:"required"`
	Password  string               `json:"password" binding:"required"`
	Phone     string               `json:"telepon" binding:"required"`
	Role      entity.SliceOfString `json:"role" binding:"required"`
	CreatedAt time.Time           `json:"createdAt"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
