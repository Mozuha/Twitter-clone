package controller

import (
	"api/models"
	"api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers() []models.User
	GetUser(ctx *gin.Context) models.User
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

func (c *controller) GetUser(ctx *gin.Context) models.User {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	return c.service.GetUser(uint(id))
}
