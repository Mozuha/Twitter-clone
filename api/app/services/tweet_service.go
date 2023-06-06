package services

import (
	"app/ent"
	"context"
	"errors"
)

type tweetService struct {
	client *ent.Client
}

// Also handles getById and those kinds by specifying 'where' argument
func (t *tweetService) GetTweets(ctx context.Context, where *ent.TweetWhereInput) ([]*ent.Tweet, error) {
	var (
		tweets []*ent.Tweet
		err    error
	)

	pred, err := where.P()
	if err != nil {
		if err.Error() == "ent: empty predicate TweetWhereInput" {
			// for getting all tweets (no where predicate)
			tweets, err = t.client.Tweet.Query().All(ctx)
		} else {
			return nil, err
		}
	} else {
		tweets, err = t.client.Tweet.Query().Where(pred).All(ctx)

		// even when no record was matched, All() will return empty slice and deem it not as an error
		// need to set not found error if no record was matched
		if len(tweets) == 0 {
			err = errors.New("ent: tweet not found")
		}
	}

	if err != nil {
		return nil, err
	}

	return tweets, nil
}

func (t *tweetService) CreateTweet(ctx context.Context, input ent.CreateTweetInput) (*ent.Tweet, error) {
	tweet, err := t.client.Tweet.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (t *tweetService) DeleteTweetById(ctx context.Context, id int) (*bool, error) {
	err := t.client.Tweet.DeleteOneID(id).Exec(ctx)
	isOk := err == nil
	if !isOk {
		return &isOk, err
	}

	return &isOk, err
}
