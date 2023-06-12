package services

import (
	"app"
	"app/auth"
	"app/ent"
	"app/ent/user"
	"app/utils"
	"context"
	"strconv"

	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"
)

type signinService struct {
	client *ent.Client
}

func (s *signinService) Signin(ctx context.Context, email string, password string) (*app.SigninResponse, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.NOT_FOUND, "no user with given email")
		return nil, gErr
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.UNAUTHORIZED, "password incorrect")
		return nil, gErr
	}

	uId := strconv.Itoa(user.ID)

	session := sessions.Default(gc)
	session.Set("user", uId)
	session.Options(sessions.Options{MaxAge: 86400 * 14})
	if err = session.Save(); err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "")
		return nil, gErr
	}

	sId := session.ID()

	accToken, err := auth.GenerateToken(uId, sId, true)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to generate access token")
		return nil, gErr
	}

	refToken, err := auth.GenerateToken(uId, sId, false)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to generate refresh token")
		return nil, gErr
	}

	return &app.SigninResponse{UserID: user.ID, AccessToken: accToken, RefreshToken: refToken}, nil
}

func (s *signinService) Signout(ctx context.Context) (*bool, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "")
		return nil, gErr
	}

	session := sessions.Default(gc)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	err = session.Save()

	isOk := err == nil
	if !isOk {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "")
		return &isOk, gErr
	}

	return &isOk, nil
}

func (s *signinService) RefreshToken(ctx context.Context, refTokenString string) (string, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "")
		return "", gErr
	}

	sessionId := sessions.Default(gc).ID()
	token, err := auth.RefreshToken(sessionId, refTokenString)
	if err != nil {
		gErr := utils.CreateGqlErr(ctx, err, utils.INTERNAL_SERVER_ERROR, "failed to refresh token")
		return "", gErr
	}

	return token, nil
}
