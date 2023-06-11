package services

import (
	"app"
	"app/auth"
	"app/ent"
	"app/ent/user"
	"app/utils"
	"context"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-contrib/sessions"
	"github.com/vektah/gqlparser/v2/gqlerror"
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
		// TODO: change other parts using gc.JSON for error handling to use gqlerror.Error and apply changes to tests
		gErr := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "User not found",
			Extensions: map[string]interface{}{
				"code":        "NOT_FOUND",
				"userMessage": "no user with given email",
				"errorDetail": err.Error(),
			},
		}
		return nil, gErr
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// gc.JSON(http.StatusUnauthorized, gin.H{"error": "Password incorrect"})
		return nil, fmt.Errorf("password incorrect: %w", err)
	}

	uId := strconv.Itoa(user.ID)

	session := sessions.Default(gc)
	session.Set("user", uId)
	session.Options(sessions.Options{MaxAge: 86400 * 14})
	if err = session.Save(); err != nil {
		// gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return nil, err
	}

	sId := session.ID()

	accToken, err := auth.GenerateToken(uId, sId, true)
	if err != nil {
		// gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return nil, err
	}

	refToken, err := auth.GenerateToken(uId, sId, false)
	if err != nil {
		// gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return nil, err
	}

	return &app.SigninResponse{UserID: user.ID, AccessToken: accToken, RefreshToken: refToken}, nil
}

func (s *signinService) Signout(ctx context.Context) (*bool, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	session := sessions.Default(gc)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	err = session.Save()

	isOk := err == nil
	if !isOk {
		return &isOk, err
	}

	return &isOk, nil
}

func (s *signinService) RefreshToken(ctx context.Context, refTokenString string) (string, error) {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	sessionId := sessions.Default(gc).ID()

	return auth.RefreshToken(sessionId, refTokenString)
}
