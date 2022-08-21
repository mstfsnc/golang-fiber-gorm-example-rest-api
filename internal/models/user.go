package models

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

type User struct {
	Id        uint64 `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
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

type SigninRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) ToUserResponse() UserResponse {
	return UserResponse{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (req SigninRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.When(req.Email == "", validation.Required.Error("username or email is required"))),
		validation.Field(&req.Email, validation.When(req.Username == "", validation.Required.Error("username or email is required"), is.EmailFormat)),
		validation.Field(&req.Password, validation.Required),
	)
}

func (req SignupRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required),
		validation.Field(&req.Email, validation.Required, is.EmailFormat),
		validation.Field(&req.Password, validation.Required),
	)
}
