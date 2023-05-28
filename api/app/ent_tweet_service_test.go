package app

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

type EntTweetServiceTestSuite struct {
	suite.Suite
	db  *ent.Client
	ctx context.Context
}

func (s *EntTweetServiceTestSuite) SetupTest() {
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
}

func (s *EntTweetServiceTestSuite) TearDownTest() {
	s.db.Close()
}

func TestEntTweetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EntTweetServiceTestSuite))
}

func (s *EntTweetServiceTestSuite) TestGetTweets() {
	tweets, err := s.db.Tweet.Query().All(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.NotEmpty(tweets)
}

func (s *EntTweetServiceTestSuite) TestGetTweetByID() {
	targetTweet, err := s.db.Tweet.Query().First(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.Run("success", func() {
		tweet, err := s.db.Tweet.Get(s.ctx, targetTweet.ID)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(targetTweet.ID, tweet.ID)
		s.Equal(targetTweet.Edges.PostedBy, tweet.Edges.PostedBy)
	})

	s.Run("error/not found", func() {
		_, err := s.db.Tweet.Get(s.ctx, 100)
		if !ent.IsNotFound(err) {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(true, ent.IsNotFound(err))
	})
}

func (s *EntTweetServiceTestSuite) TestCreateTweet() {
	user, err := s.db.User.Query().First(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}
	expectedTweet := &ent.Tweet{
		Text:  "sample tweet2",
		Edges: ent.TweetEdges{PostedBy: user},
	}

	s.Run("success", func() {
		tweet, err := s.db.Tweet.
			Create().
			SetText(expectedTweet.Text).
			SetPostedBy(expectedTweet.Edges.PostedBy).
			Save(s.ctx)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}
		postedBy, err := tweet.PostedBy(s.ctx)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(expectedTweet.Text, tweet.Text)
		s.Equal(expectedTweet.Edges.PostedBy.ID, postedBy.ID)
	})

	s.Run("error/text field (required) is missing", func() {
		err := s.db.Tweet.
			Create().
			SetPostedBy(expectedTweet.Edges.PostedBy).
			Exec(s.ctx)

		s.Equal(true, s.Error(err))
	})

	// TODO: Text length check
}

func (s *EntTweetServiceTestSuite) TestDeleteTweetByID() {
	targetTweet, err := s.db.Tweet.Query().Where(tweet.HasPostedByWith(user.Email("test3@gmail.com"))).Only(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.Run("success", func() {
		err := s.db.Tweet.DeleteOneID(targetTweet.ID).Exec(s.ctx)
		s.NoError(err)
		_, err = s.db.Tweet.Query().Where(tweet.ID(targetTweet.ID)).Only(s.ctx)
		s.Equal(true, ent.IsNotFound(err))
	})

	s.Run("error/not found", func() {
		err := s.db.Tweet.DeleteOneID(100).Exec(s.ctx)
		s.Equal(true, ent.IsNotFound(err))
	})
}
