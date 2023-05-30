package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"app/ent"
	"app/gql/generated"
	"context"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return r.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*bool, error) {
	err := r.client.User.DeleteOneID(id).Exec(ctx)
	isErr := err == nil
	return &isErr, err
}

// CreateTweet is the resolver for the createTweet field.
func (r *mutationResolver) CreateTweet(ctx context.Context, input ent.CreateTweetInput) (*ent.Tweet, error) {
	return r.client.Tweet.Create().SetInput(input).Save(ctx)
}

// DeleteTweet is the resolver for the deleteTweet field.
func (r *mutationResolver) DeleteTweet(ctx context.Context, id int) (*bool, error) {
	err := r.client.Tweet.DeleteOneID(id).Exec(ctx)
	isErr := err == nil
	return &isErr, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
