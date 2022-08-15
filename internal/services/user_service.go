package services

import (
	"sample-app/internal/models"
	"sample-app/internal/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) All() ([]models.User, error) {
	return s.userRepository.All()
}

func (s UserService) RetrieveById(id uint64) (models.User, error) {
	return s.userRepository.Retrieve(id)
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
