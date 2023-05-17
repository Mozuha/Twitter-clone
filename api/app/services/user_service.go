package services

import (
	"api/models"
	"errors"
)

type UserService interface {
	GetUsers() []models.User
	GetUserById(uint) (models.User, error)
	CreateUser(models.User)
}

type userService struct {
	users []models.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) GetUsers() []models.User {
	return service.users
}

func (service *userService) GetUserById(id uint) (models.User, error) {
	for _, c := range service.users {
		if c.Id == id {
			return c, nil
		}
	}
	return service.users[1], errors.New("user with specified id not found")
}

func (service *userService) CreateUser(user models.User) {
	service.users = append(service.users, user)
}
