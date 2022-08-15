package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	Id        uint64    `gorm:"primaryKey"`
	Username  string    `json:"username" validate:"required,min=3,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserResponse struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func (u User) Validate() []*ValidationError {
	var validate = validator.New()
	var errors []*ValidationError
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
			})
		}
	}
	return errors
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
