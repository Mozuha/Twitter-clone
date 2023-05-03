package services

import "api/models"

type UserService interface {
	GetUsers() []models.User
	GetUser(uint) models.User
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

func (service *userService) GetUser(id uint) models.User {
	for _, c := range service.users {
		if c.Id == id {
			return c
		}
	}
	// TODO: what if user with specified id was not found? return nil but how? search
	// *for now, assuming that user with specified id is always exist
	return service.users[id]
}
