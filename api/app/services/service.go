package services

import (
	"app"
	"app/ent"
	"context"
)

type Services interface {
	UserService
	TweetService
	SigninService
	NodeService
}

type services struct {
	*userService
	*tweetService
	*signinService
	*nodeService
}

func New(client *ent.Client) Services {
	return &services{
		userService:   &userService{client: client},
		tweetService:  &tweetService{client: client},
		signinService: &signinService{client: client},
		nodeService:   &nodeService{client: client},
	}
}

type UserService interface {
	GetUsers(ctx context.Context, conn *UsersConnection) (*ent.UserConnection, error)
	CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error)
	UpdateUserById(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error)
	DeleteUserById(ctx context.Context, id int) (*bool, error)
	CheckEmailExists(ctx context.Context, email string) (*bool, error)
	CheckScreenNameExists(ctx context.Context, screenName string) (*bool, error)
}

type TweetService interface {
	GetTweets(ctx context.Context, conn *TweetsConnection) (*ent.TweetConnection, error)
	CreateTweet(ctx context.Context, input ent.CreateTweetInput) (*ent.Tweet, error)
	DeleteTweetById(ctx context.Context, id int) (*bool, error)
}

type SigninService interface {
	Signin(ctx context.Context, email string, password string) (*app.SigninResponse, error)
	Signout(ctx context.Context) (*bool, error)
	RefreshToken(ctx context.Context, refTokenString string) (string, error)
}

type NodeService interface {
	Node(ctx context.Context, id int) (ent.Noder, error)
	Nodes(ctx context.Context, ids []int) ([]ent.Noder, error)
}
