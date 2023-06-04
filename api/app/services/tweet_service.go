package services

import (
	"app/ent"
	"context"
	"fmt"
)

type tweetService struct {
	client *ent.Client
}

// Also handles getById and those kinds by specifying 'where' argument
func (t *tweetService) GetTweets(ctx context.Context, where *ent.TweetWhereInput) ([]*ent.Tweet, error) {
	pred, err := where.P()
	if err != nil {
		return nil, err
	}

	tweets, err := t.client.Tweet.Query().Where(pred).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting tweets: %w", err)
	}

	return tweets, nil
}

func (t *tweetService) CreateTweet(ctx context.Context, input ent.CreateTweetInput) (*ent.Tweet, error) {
	tweet, err := t.client.Tweet.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating tweet: %w", err)
	}

	return tweet, nil
}

func (t *tweetService) DeleteTweetById(ctx context.Context, id int) (*bool, error) {
	err := t.client.Tweet.DeleteOneID(id).Exec(ctx)
	isOk := err == nil
	if !isOk {
		return &isOk, fmt.Errorf("deleting tweet: %w", err)
	}

	return &isOk, err
}
