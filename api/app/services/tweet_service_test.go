package services

import (
	"app/db"
	"app/ent"
	"app/ent/tweet"
	"app/ent/user"
	"app/utils"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const TWEET_NOT_FOUND_ERROR = "input: ent: tweet not found"

type TweetServiceTestSuite struct {
	suite.Suite
	db      *ent.Client
	ctx     context.Context
	service Services
}

func (s *TweetServiceTestSuite) SetupTest() {
	runningEnv, err := utils.LoadEnv()
	if err != nil {
		s.Fail("failed to load env: ", err)
		os.Exit(2)
	}

	s.db, err = db.ConnectTestDB(runningEnv)
	if err != nil {
		s.Fail("failed to connect to db: ", err)
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

	s.Run("success", func() {
		tweetsConn, err := s.service.GetTweets(s.ctx, &TweetsConnection{Where: &ent.TweetWhereInput{}})

		s.NotEmpty(tweetsConn)
		s.NoError(err)
	})

	// assuming that there are 3 tweets in the database
	s.Run("success/pagination", func() {
		first := 2
		// take first 2 tweets
		tweetsConn, err := s.service.GetTweets(s.ctx, &TweetsConnection{First: &first, Where: &ent.TweetWhereInput{}})

		s.Equal(first, len(tweetsConn.Edges))
		s.NoError(err)

		// take the remaining 1 tweet after the last cursor of the first 2 tweet
		tweetsConn, err = s.service.GetTweets(s.ctx, &TweetsConnection{After: tweetsConn.PageInfo.EndCursor, Where: &ent.TweetWhereInput{}})

		s.Equal(1, len(tweetsConn.Edges))
		s.NoError(err)
	})
}

func (s *TweetServiceTestSuite) TestGetTweetByID() {
	targetTweet, err := s.db.Tweet.Query().First(s.ctx)
	if err != nil {
		s.Fail("failed to get tweet to be used: ", err)
	}

	s.Run("success", func() {
		tweetConn, err := s.service.GetTweets(s.ctx, &TweetsConnection{Where: &ent.TweetWhereInput{ID: &targetTweet.ID}})

		s.Equal(targetTweet.ID, tweetConn.Edges[0].Node.ID)
		s.Equal(targetTweet.Edges.PostedBy, tweetConn.Edges[0].Node.Edges.PostedBy)
		s.NoError(err)
	})

	s.Run("error/not found", func() {
		notExistingId := 100
		_, err := s.service.GetTweets(s.ctx, &TweetsConnection{Where: &ent.TweetWhereInput{ID: &notExistingId}})

		if err.Error() != TWEET_NOT_FOUND_ERROR {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(true, err.Error() == TWEET_NOT_FOUND_ERROR)
	})
}

func (s *TweetServiceTestSuite) TestCreateTweet() {
	postUser, err := s.db.User.Query().First(s.ctx)
	if err != nil {
		s.Fail("failed to get user to be used: ", err)
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
		s.NoError(err)

		postedBy, err := tweet.PostedBy(s.ctx)
		s.NoError(err)

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
		s.Fail("failed to get tweet to be deleted: ", err)
	}

	s.Run("success", func() {
		isDeleted, err := s.service.DeleteTweetById(s.ctx, targetTweet.ID)
		s.Equal(true, *isDeleted)
		s.NoError(err)

		_, err = s.service.GetTweets(s.ctx, &TweetsConnection{Where: &ent.TweetWhereInput{ID: &targetTweet.ID}})
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
