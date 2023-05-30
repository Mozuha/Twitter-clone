package db

import (
	"app/ent"
	"context"
	"fmt"
)

func InsertMockData(ctx context.Context, client *ent.Client) error {

	// Refresh table; Be aware of the order of deletion as Tweets are depending on Users
	numDel, err := client.Tweet.Delete().Exec(ctx)
	if err != nil {
		return fmt.Errorf("deleting existing mock tweets: %w", err)
	}
	fmt.Printf("deleted %d existing mock tweets\n", numDel)

	numDel, err = client.User.Delete().Exec(ctx)
	if err != nil {
		return fmt.Errorf("deleting existing mock users: %w", err)
	}
	fmt.Printf("deleted %d existing mock users\n", numDel)

	// Users
	usersMock := []ent.User{
		{ScreenName: "test1", Name: "test 1", Email: "test1@gmail.com", Password: "pass", ProfileImage: "images/test1.jpg"},
		{ScreenName: "test2", Name: "test 2", Email: "test2@ymail.ne.jp", Password: "word", ProfileImage: "images/test2.jpg"},
		{ScreenName: "test3", Name: "test 3", Email: "test3@gmail.com", Password: "password", ProfileImage: "images/test3.jpg"},
		{ScreenName: "test4", Name: "test 4", Email: "test4@ymail.ne.jp", Password: "tobedeleted", ProfileImage: "images/test4.jpg"},
	}

	usersBulk := make([]*ent.UserCreate, len(usersMock))
	for i, user := range usersMock {
		usersBulk[i] = client.User.
			Create().
			SetName(user.Name).
			SetScreenName(user.ScreenName).
			SetEmail(user.Email).
			SetPassword(user.Password).
			SetProfileImage(user.ProfileImage)
	}

	users, err := client.User.CreateBulk(usersBulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("creating mock users: %w\n", err)
	}
	fmt.Println("created mock users: ", users)

	// Tweets
	tweetsMock := []ent.Tweet{
		{Text: "sample tweet", Edges: ent.TweetEdges{PostedBy: users[0]}},
		{Text: "lorem ipsum", Edges: ent.TweetEdges{PostedBy: users[1]}},
		{Text: "逍遙遊", Edges: ent.TweetEdges{PostedBy: users[2]}},
	}

	tweetsBulk := make([]*ent.TweetCreate, len(tweetsMock))
	for i, tweet := range tweetsMock {
		tweetsBulk[i] = client.Tweet.
			Create().
			SetText(tweet.Text).
			SetPostedBy(tweet.Edges.PostedBy)
	}

	tweets, err := client.Tweet.CreateBulk(tweetsBulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("creating mock tweets: %w\n", err)
	}
	fmt.Println("created mock tweets: ", tweets)

	return nil
}
