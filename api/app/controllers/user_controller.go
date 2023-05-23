package controllers

import (
	"app/models"
	"app/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(*gin.Context) ([]models.User, error)
	GetUserById(*gin.Context) (models.User, error)
	CreateUser(*gin.Context) (models.User, error)
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) GetUsers(ctx *gin.Context) ([]models.User, error) {
	return c.service.GetUsers(ctx)
}

func (c *userController) GetUserById(ctx *gin.Context) (models.User, error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	return c.service.GetUserById(ctx, uint(id))
}

func (c *userController) CreateUser(ctx *gin.Context) (models.User, error) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return user, err
	}

	user.ProfileImage = "images/default.jpg"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return c.service.CreateUser(ctx, user)
}
