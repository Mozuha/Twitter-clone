package services

import (
	"app/auth"
	"app/db"
	"app/ent"
	"app/utils"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SigninServiceTestSuite struct {
	suite.Suite
	db      *ent.Client
	ctx     context.Context
	service Services
}

func (s *SigninServiceTestSuite) SetupTest() {
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

func (s *SigninServiceTestSuite) TearDownTest() {
	s.db.Close()
}

func TestSigninServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SigninServiceTestSuite))
}

func (s *SigninServiceTestSuite) TestSignin() {
	input := ent.CreateUserInput{
		Name:       "test 5",
		ScreenName: "test5",
		Email:      "test5@ymail.ne.jp",
		Password:   "12345",
	}

	targetUser, err := s.service.CreateUser(s.ctx, input)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.Run("success", func() {
		res, err := s.service.Signin(s.ctx, targetUser.Email, "12345")
		accToken, err := auth.ValidateToken(res.AccessToken)
		refToken, err := auth.ValidateToken(res.RefreshToken)

		s.Equal(targetUser.ID, res.UserID)
		s.NotEmpty(accToken)
		s.NotEmpty(refToken)
		s.NoError(err)
	})

	s.Run("error/email incorrect", func() {
		res, err := s.service.Signin(s.ctx, "tst5@ymail.ne.jp", "12345")

		s.Empty(res)
		s.ErrorContains(err, "no user with given email")
	})

	s.Run("error/password incorrect", func() {
		res, err := s.service.Signin(s.ctx, "test5@ymail.ne.jp", "passssssss")

		s.Empty(res)
		s.ErrorContains(err, "password incorrect")
	})
}
