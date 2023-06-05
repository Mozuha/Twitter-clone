package services

import (
	"app"
	"app/ent"
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type Services interface {
	UserService
	TweetService
	SigninService
	JWTService
	NodeService
}

type services struct {
	*userService
	*tweetService
	*signinService
	*jwtService
	*nodeService
}

func New(client *ent.Client) Services {
	return &services{
		userService:   &userService{client: client},
		tweetService:  &tweetService{client: client},
		signinService: &signinService{client: client},
		jwtService:    &jwtService{issuer: "example_issuer"},
		nodeService:   &nodeService{client: client},
	}
}

type UserService interface {
	GetUsers(ctx context.Context, where *ent.UserWhereInput) ([]*ent.User, error)
	CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error)
	UpdateUserById(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error)
	DeleteUserById(ctx context.Context, id int) (*bool, error)
	CheckEmailExists(ctx context.Context, email string) (*bool, error)
	CheckScreenNameExists(ctx context.Context, screenName string) (*bool, error)
}

type TweetService interface {
	GetTweets(ctx context.Context, where *ent.TweetWhereInput) ([]*ent.Tweet, error)
	CreateTweet(ctx context.Context, input ent.CreateTweetInput) (*ent.Tweet, error)
	DeleteTweetById(ctx context.Context, id int) (*bool, error)
}

type SigninService interface {
	Signin(ctx context.Context, email string, password string) (*app.SigninResponse, error)
}

type JWTService interface {
	GenerateToken(screenName string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type NodeService interface {
	Node(ctx context.Context, id int) (ent.Noder, error)
	Nodes(ctx context.Context, ids []int) ([]ent.Noder, error)
}
