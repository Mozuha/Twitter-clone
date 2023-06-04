package services

import (
	"app/ent"
	"app/ent/user"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	client *ent.Client
}

// Also handles getById and getByEmail and those kinds by specifying 'where' argument
func (u *userService) GetUsers(ctx context.Context, where *ent.UserWhereInput) ([]*ent.User, error) {
	pred, err := where.P()
	if err != nil {
		return nil, err
	}

	users, err := u.client.User.Query().Where(pred).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting users: %w", err)
	}

	return users, nil
}

func (u *userService) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	input.Password = string(hash)
	user, err := u.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	return user, nil
}

func (u *userService) UpdateUserById(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	user, err := u.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("updating user: %w", err)
	}

	return user, nil
}

func (u *userService) DeleteUserById(ctx context.Context, id int) (*bool, error) {
	err := u.client.User.DeleteOneID(id).Exec(ctx)
	isOk := err == nil
	if !isOk {
		return &isOk, fmt.Errorf("deleting user: %w", err)
	}

	return &isOk, nil
}

func (u *userService) CheckEmailExists(ctx context.Context, email string) (*bool, error) {
	user, err := u.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	isEmailExists := user != nil
	if !isEmailExists {
		return &isEmailExists, err
	}

	return &isEmailExists, nil
}

func (u *userService) CheckScreenNameExists(ctx context.Context, screenName string) (*bool, error) {
	user, err := u.client.User.Query().Where(user.ScreenNameEQ(screenName)).Only(ctx)
	isScreenNameExists := user != nil
	if !isScreenNameExists {
		return &isScreenNameExists, err
	}

	return &isScreenNameExists, nil
}
