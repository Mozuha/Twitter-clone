package services

import (
	"app/ent"
	"app/models"
	"context"
	"fmt"
)

// TODO: create and use response type instead of models.User type
type UserService interface {
	GetUsers(context.Context) ([]models.User, error)
	GetUserById(context.Context, uint) (models.User, error)
	CreateUser(context.Context, models.User) (models.User, error)
}

type userService struct {
	client *ent.UserClient
}

func NewUserService(userClient *ent.UserClient) UserService {
	return &userService{
		client: userClient,
	}
}

func (srv *userService) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := srv.client.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting users: %w", err)
	}

	var res []models.User
	for _, user := range users {
		res = append(res, models.User{
			Id:           uint(user.ID),
			ScreenName:   user.ScreenName,
			Username:     user.Name,
			Email:        user.Email,
			Password:     user.Password,
			ProfileImage: user.ProfileImage,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		})
	}

	return res, nil
}

func (srv *userService) GetUserById(ctx context.Context, id uint) (models.User, error) {
	user, err := srv.client.Get(ctx, int(id))
	if err != nil {
		return models.User{}, fmt.Errorf("getting user: %w", err)
	}

	res := models.User{
		Id:           uint(user.ID),
		ScreenName:   user.ScreenName,
		Username:     user.Name,
		Email:        user.Email,
		Password:     user.Password,
		ProfileImage: user.ProfileImage,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return res, nil
}

func (srv *userService) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	newUser, err := srv.client.
		Create().
		SetScreenName(user.ScreenName).
		SetName(user.Username).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetProfileImage(user.ProfileImage).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(ctx)
	if err != nil {
		return models.User{}, fmt.Errorf("creating user: %w", err)
	}

	res := models.User{
		Id:           uint(newUser.ID),
		ScreenName:   newUser.ScreenName,
		Username:     newUser.Name,
		Email:        newUser.Email,
		Password:     newUser.Password,
		ProfileImage: newUser.ProfileImage,
		CreatedAt:    newUser.CreatedAt,
		UpdatedAt:    newUser.UpdatedAt,
	}

	return res, nil
}
