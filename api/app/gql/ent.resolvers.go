package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"app/ent"
	"app/gql/generated"
	"context"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Tweets is the resolver for the tweets field.
func (r *queryResolver) Tweets(ctx context.Context, where *ent.TweetWhereInput) ([]*ent.Tweet, error) {
	pred, err := where.P()
	if err != nil {
		return nil, err
	}
	return r.client.Tweet.Query().Where(pred).All(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, where *ent.UserWhereInput) ([]*ent.User, error) {
	pred, err := where.P()
	if err != nil {
		return nil, err
	}
	return r.client.User.Query().Where(pred).All(ctx)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
