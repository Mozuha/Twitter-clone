package handlers

import (
	"app/controllers"
	"app/ent"
	"app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler interface {
	GetUsersHandler(*gin.Context)
	GetUserByIdHandler(*gin.Context)
	CreateUserHandler(*gin.Context)
}

type usersHandler struct {
	controller controllers.UserController
}

func NewUsersHandler(client *ent.Client) UsersHandler {
	userService := services.NewUserService(client.User)
	userController := controllers.NewUserController(userService)
	return &usersHandler{
		controller: userController,
	}
}

func (h *usersHandler) GetUsersHandler(ctx *gin.Context) {
	users, err := h.controller.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (h *usersHandler) GetUserByIdHandler(ctx *gin.Context) {
	user, err := h.controller.GetUserById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func (h *usersHandler) CreateUserHandler(ctx *gin.Context) {
	user, err := h.controller.CreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, user)
	}
}
