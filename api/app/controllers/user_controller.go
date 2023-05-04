package controllers

import (
	"api/models"
	"api/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers() []models.User
	GetUserById(ctx *gin.Context) (models.User, error)
	CreateUser(ctx *gin.Context) (models.User, error)
}

type controller struct {
	service services.UserService
}

func New(service services.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetUsers() []models.User {
	return c.service.GetUsers()
}

func (c *controller) GetUserById(ctx *gin.Context) (models.User, error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	user, err := c.service.GetUserById(uint(id))
	return user, err
}

func (c *controller) CreateUser(ctx *gin.Context) (models.User, error) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return user, err
	}
	user.Id = 4 // id creation will be handled by db; here, use arbitrary value
	user.ProfileImage = "images/default.jpg"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	c.service.CreateUser(user)
	return user, nil
}
