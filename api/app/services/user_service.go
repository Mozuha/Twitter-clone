package services

import (
	"app/ent"
	"app/ent/user"
	"app/utils"
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	client *ent.Client
}

type UsersConnection struct {
	After   *entgql.Cursor[int]
	First   *int
	Before  *entgql.Cursor[int]
	Last    *int
	OrderBy *ent.UserOrder
	Where   *ent.UserWhereInput
}

// Also handles getById and getByEmail and those kinds by specifying 'where' argument
func (u *userService) GetUsers(ctx context.Context, conn *UsersConnection) (*ent.UserConnection, error) {
	var (
		usersConn *ent.UserConnection
		err       error
	)

	_, err = conn.Where.P()
	if err != nil {
		if err.Error() == "ent: empty predicate UserWhereInput" {
			// for getting all users (no where predicate)
			usersConn, err = u.client.User.Query().Paginate(ctx, conn.After, conn.First, conn.Before, conn.Last, ent.WithUserOrder(conn.OrderBy))
		} else {
			gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to parse where predicate")
			return nil, gErr
		}
	} else {
		usersConn, err = u.client.User.Query().Paginate(ctx, conn.After, conn.First, conn.Before, conn.Last, ent.WithUserOrder(conn.OrderBy), ent.WithUserFilter(conn.Where.Filter))

		// even when no record was matched, All() will return empty slice and deem it not as an error
		// need to set not found error if no record was matched
		if usersConn.TotalCount == 0 {
			err = errors.New("ent: user not found")
		}
	}

	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to get users")
		return nil, gErr
	}

	return usersConn, nil
}

func (u *userService) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "")
		return nil, gErr
	}

	input.Password = string(hash)
	user, err := u.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to create user")
		return nil, gErr
	}

	return user, nil
}

func (u *userService) UpdateUserById(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	// updated password must also be hashed
	if input.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "")
			return nil, gErr
		}
		hashStr := string(hash)
		input.Password = &hashStr
	}

	user, err := u.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to update user")
		return nil, gErr
	}

	return user, nil
}

func (u *userService) DeleteUserById(ctx context.Context, id int) (*bool, error) {
	err := u.client.User.DeleteOneID(id).Exec(ctx)
	isOk := err == nil
	if !isOk {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to delete user")
		return &isOk, gErr
	}

	// TODO: should the session be deleted here?

	return &isOk, nil
}

func (u *userService) CheckEmailExists(ctx context.Context, email string) (*bool, error) {
	user, err := u.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	isEmailExists := user != nil
	if !isEmailExists {
		if ent.IsNotFound(err) {
			return &isEmailExists, nil // false, i.e. not exists
		} else {
			return nil, err // other kind of error
		}
	}

	return &isEmailExists, nil // true, i.e. exists
}

func (u *userService) CheckScreenNameExists(ctx context.Context, screenName string) (*bool, error) {
	user, err := u.client.User.Query().Where(user.ScreenNameEQ(screenName)).Only(ctx)
	isScreenNameExists := user != nil
	if !isScreenNameExists {
		if ent.IsNotFound(err) {
			return &isScreenNameExists, err
		} else {
			return nil, err
		}
	}

	return &isScreenNameExists, nil
}
