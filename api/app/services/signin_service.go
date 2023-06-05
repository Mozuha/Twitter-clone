package services

import (
	"app"
	"app/ent"
	"app/ent/user"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type signinService struct {
	client *ent.Client
}

func (s *signinService) Signin(ctx context.Context, email string, password string) (*app.SigninResponse, error) {
	user, err := s.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("no user with given email: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("password incorrect: %w", err)
	}

	token, err := New(s.client).GenerateToken(user.ScreenName)
	if err != nil {
		return nil, err
	}

	return &app.SigninResponse{UserID: user.ID, Token: token}, nil
}
