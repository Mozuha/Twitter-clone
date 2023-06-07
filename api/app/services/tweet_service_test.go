package services

import (
	"app/db"
	"app/ent"
	"app/ent/tweet"
	"app/ent/user"
	"app/utils"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const TWEET_NOT_FOUND_ERROR = "ent: tweet not found"

type TweetServiceTestSuite struct {
	suite.Suite
	db      *ent.Client
	ctx     context.Context
	service Services
}

func (s *TweetServiceTestSuite) SetupTest() {
	runningEnv, err := utils.LoadEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	s.db, err = db.ConnectTestDB(runningEnv)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	s.ctx = context.Background()

	s.service = New(s.db)
}

func (s *TweetServiceTestSuite) TearDownTest() {
	s.db.Close()
}

func TestTweetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TweetServiceTestSuite))
}

func (s *TweetServiceTestSuite) TestGetTweets() {
	tweets, err := s.service.GetTweets(s.ctx, &ent.TweetWhereInput{})
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.NotEmpty(tweets)
}

func (s *TweetServiceTestSuite) TestGetTweetByID() {
	targetTweet, err := s.db.Tweet.Query().First(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.Run("success", func() {
		tweet, err := s.service.GetTweets(s.ctx, &ent.TweetWhereInput{ID: &targetTweet.ID})
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(targetTweet.ID, tweet[0].ID)
		s.Equal(targetTweet.Edges.PostedBy, tweet[0].Edges.PostedBy)
	})

	s.Run("error/not found", func() {
		notExistingId := 100
		_, err := s.service.GetTweets(s.ctx, &ent.TweetWhereInput{ID: &notExistingId})

		if err.Error() != TWEET_NOT_FOUND_ERROR {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(true, err.Error() == TWEET_NOT_FOUND_ERROR)
	})
}

func (s *TweetServiceTestSuite) TestCreateTweet() {
	postUser, err := s.db.User.Query().First(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}
	expectedTweet := &ent.Tweet{
		Text:  "sample tweet2",
		Edges: ent.TweetEdges{PostedBy: postUser},
	}

	s.Run("success", func() {
		input := ent.CreateTweetInput{
			Text:       expectedTweet.Text,
			PostedByID: postUser.ID,
		}

		tweet, err := s.service.CreateTweet(s.ctx, input)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		postedBy, err := tweet.PostedBy(s.ctx)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(expectedTweet.Text, tweet.Text)
		s.Equal(postUser.ID, postedBy.ID)
	})

	s.Run("error/text field (required) is missing", func() {
		input := ent.CreateTweetInput{
			PostedByID: postUser.ID,
		}

		_, err := s.service.CreateTweet(s.ctx, input)

		s.Error(err)
	})

	// TODO: Text length check
}

func (s *TweetServiceTestSuite) TestDeleteTweetByID() {
	targetTweet, err := s.db.Tweet.Query().Where(tweet.HasPostedByWith(user.Email("test3@gmail.com"))).Only(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.Run("success", func() {
		isDeleted, err := s.service.DeleteTweetById(s.ctx, targetTweet.ID)
		s.Equal(true, *isDeleted)
		s.NoError(err)

		_, err = s.db.Tweet.Query().Where(tweet.ID(targetTweet.ID)).Only(s.ctx)
		s.Equal(true, err.Error() == TWEET_NOT_FOUND_ERROR)
	})

	s.Run("error/not found", func() {
		_, err := s.service.DeleteTweetById(s.ctx, 100)
		s.Equal(true, err.Error() == TWEET_NOT_FOUND_ERROR)
	})
}

// TODO: fetch tweet's likes
// TODO: fetch tweet's replies
// TODO: fetch tweet's parent
// TODO: fetch tweet's owner
