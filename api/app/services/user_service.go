package services

import (
	"api/models"
	"errors"
	"time"
)

type UserService interface {
	GetUsers() []models.User
	GetUserById(uint) (models.User, error)
	CreateUser(models.User)
}

type userService struct {
	users []models.User
}

// for mockup
var usersMock = []models.User{
	{Id: 1, DisplayName: "test1", Username: "test 1", Email: "test1@gmail.com", Password: "pass", ProfileImage: "images/test1.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Id: 2, DisplayName: "test2", Username: "test 2", Email: "test2@ymail.ne.jp", Password: "word", ProfileImage: "images/test2.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Id: 3, DisplayName: "test3", Username: "test 3", Email: "test3@gmail.com", Password: "password", ProfileImage: "images/test3.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func New() UserService {
	return &userService{usersMock}
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
