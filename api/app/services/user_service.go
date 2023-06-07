package services

import (
	"app/ent"
	"app/ent/user"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	client *ent.Client
}

// Also handles getById and getByEmail and those kinds by specifying 'where' argument
func (u *userService) GetUsers(ctx context.Context, where *ent.UserWhereInput) ([]*ent.User, error) {
	var (
		users []*ent.User
		err   error
	)

	pred, err := where.P()
	if err != nil {
		if err.Error() == "ent: empty predicate UserWhereInput" {
			// for getting all users (no where predicate)
			users, err = u.client.User.Query().All(ctx)
		} else {
			return nil, err
		}
	} else {
		users, err = u.client.User.Query().Where(pred).All(ctx)

		// even when no record was matched, All() will return empty slice and deem it not as an error
		// need to set not found error if no record was matched
		if len(users) == 0 {
			err = errors.New("ent: user not found")
		}
	}

	if err != nil {
		return nil, err
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
		return nil, err
	}

	return user, nil
}

func (u *userService) UpdateUserById(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	// updated password must also be hashed
	if input.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		hashStr := string(hash)
		input.Password = &hashStr
	}

	user, err := u.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteUserById(ctx context.Context, id int) (*bool, error) {
	err := u.client.User.DeleteOneID(id).Exec(ctx)
	isOk := err == nil
	if !isOk {
		return &isOk, err
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
