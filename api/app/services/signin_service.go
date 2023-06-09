package services

import (
	"app"
	"app/auth"
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

	accToken, err := auth.GenerateToken(user.ScreenName, true)
	if err != nil {
		return nil, err
	}

	refToken, err := auth.GenerateToken(user.ScreenName, false)
	if err != nil {
		return nil, err
	}

	return &app.SigninResponse{UserID: user.ID, AccessToken: accToken, RefreshToken: refToken}, nil
}
