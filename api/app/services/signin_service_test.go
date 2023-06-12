package services

import (
	"app/auth"
	"app/db"
	"app/ent"
	"app/utils"
	"context"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type SigninServiceTestSuite struct {
	suite.Suite
	db      *ent.Client
	store   redis.Store
	ctx     context.Context
	service Services
}

func (s *SigninServiceTestSuite) SetupTest() {
	runningEnv, err := utils.LoadEnv()
	if err != nil {
		s.Fail("failed loading env value: ", err)
		os.Exit(2)
	}

	s.db, err = db.ConnectTestDB(runningEnv)
	if err != nil {
		s.Fail("failed connecting to test db: ", err)
		os.Exit(2)
	}

	s.store, err = db.SetUpRedisStore(runningEnv)
	if err != nil {
		s.Fail("failed setting up redis store: ", err)
		os.Exit(2)
	}

	gc, _ := gin.CreateTestContext(httptest.NewRecorder())
	gc.Request = httptest.NewRequest("GET", "/", nil)

	// set session instance to gin context
	ghf := sessions.Sessions("mysession", s.store)
	ghf(gc)

	// set gin context to context
	s.ctx = context.WithValue(gc.Request.Context(), "GinContextKey", gc)

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
		s.Fail("failed to create user: ", err)
	}

	s.Run("success", func() {
		res, err := s.service.Signin(s.ctx, targetUser.Email, input.Password)
		accToken, err := auth.ValidateToken(res.AccessToken)
		refToken, err := auth.ValidateToken(res.RefreshToken)

		s.Equal(targetUser.ID, res.UserID)
		s.NotEmpty(accToken)
		s.NotEmpty(refToken)
		s.NoError(err)
	})

	s.Run("error/email incorrect", func() {
		res, err := s.service.Signin(s.ctx, "tst5@ymail.ne.jp", input.Password)

		s.Empty(res)
		s.Equal(true, err.Error() == "input: ent: user not found")
	})

	s.Run("error/password incorrect", func() {
		res, err := s.service.Signin(s.ctx, targetUser.Email, "passssssss")

		s.Empty(res)
		s.Equal(true, err.Error() == "input: crypto/bcrypt: hashedPassword is not the hash of the given password")
	})

	s.service.Signout(s.ctx)
}

func (s *SigninServiceTestSuite) TestSignout() {
	input := ent.CreateUserInput{
		Name:       "test 5",
		ScreenName: "test5",
		Email:      "test5@ymail.ne.jp",
		Password:   "12345",
	}

	targetUser, err := s.service.CreateUser(s.ctx, input)
	if err != nil {
		s.Fail("failed to create user: ", err)
	}

	s.Run("success", func() {
		_, err := s.service.Signin(s.ctx, targetUser.Email, input.Password)
		if err != nil {
			s.Fail("failed to signin: ", err)
		}

		isSignedout, err := s.service.Signout(s.ctx)
		s.Equal(true, *isSignedout)
		s.NoError(err)
	})
}
