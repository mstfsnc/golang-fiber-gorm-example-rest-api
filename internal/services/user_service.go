package services

import (
	"sample-app/internal/models"
	"sample-app/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s UserService) Retrieve(user *models.User) error {
	return s.userRepository.Retrieve(user)
}

func (s UserService) Create(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s UserService) Update(user *models.User, fields models.User) error {
	return s.userRepository.Update(user, fields)
}

func (s UserService) Delete(id uint64) error {
	return s.userRepository.Delete(id)
}
