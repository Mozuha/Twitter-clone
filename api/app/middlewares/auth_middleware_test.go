package middlewares

import (
	"app/db"
	"app/ent"
	"app/services"
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

type AuthMiddlewareTestSuite struct {
	suite.Suite
	db      *ent.Client
	store   redis.Store
	service services.Services
}

func (s *AuthMiddlewareTestSuite) SetupTest() {
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

	s.service = services.New(s.db)
}

func TestAuthMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}

func (s *AuthMiddlewareTestSuite) TestAuthMiddleware() {
	input := ent.CreateUserInput{
		Name:       "test 5",
		ScreenName: "test5",
		Email:      "test5@ymail.ne.jp",
		Password:   "12345",
	}

	targetUser, err := s.service.CreateUser(context.Background(), input)
	if err != nil {
		s.Fail("failed to create user: ", err)
	}

	s.Run("success", func() {
		gc, _ := gin.CreateTestContext(httptest.NewRecorder())
		gc.Request = httptest.NewRequest("GET", "/", nil)
		sessions.Sessions("mysession", s.store)(gc)
		c := context.WithValue(gc.Request.Context(), "GinContextKey", gc)

		res, err := s.service.Signin(c, targetUser.Email, input.Password)
		if err != nil {
			s.Fail("failed to signin: ", err)
		}

		gc.Request.Header.Add("Authorization", "Bearer "+res.AccessToken)

		AuthMiddleware()(gc)
		isAuthed := gc.Value(authedCheckCtxKey.name).(bool)

		s.Equal(true, isAuthed)

		s.service.Signout(c)
	})

	s.Run("error/no authorization header", func() {
		gc, _ := gin.CreateTestContext(httptest.NewRecorder())
		gc.Request = httptest.NewRequest("GET", "/", nil)

		AuthMiddleware()(gc)
		isAuthed := gc.Value(authedCheckCtxKey.name).(bool)

		s.Equal(false, isAuthed)
		s.Equal(true, gc.Errors.Last().Error() == "Authorization header is required")
	})

	s.Run("error/invalid token", func() {
		gc, _ := gin.CreateTestContext(httptest.NewRecorder())
		gc.Request = httptest.NewRequest("GET", "/", nil)

		invalidtkn := "invalidtoken"
		gc.Request.Header.Add("Authorization", "Bearer "+invalidtkn)

		AuthMiddleware()(gc)
		isAuthed := gc.Value(authedCheckCtxKey.name).(bool)

		s.Equal(false, isAuthed)
		s.ErrorContains(gc.Errors.Last(), "Invalid token: ")
	})

	s.Run("error/not signed in", func() {
		gc, _ := gin.CreateTestContext(httptest.NewRecorder())
		gc.Request = httptest.NewRequest("GET", "/", nil)
		sessions.Sessions("mysession", s.store)(gc)
		c := context.WithValue(gc.Request.Context(), "GinContextKey", gc)

		res, err := s.service.Signin(c, targetUser.Email, input.Password)
		if err != nil {
			s.Fail("failed to signin: ", err)
		}

		s.service.Signout(c)

		gc.Request.Header.Add("Authorization", "Bearer "+res.AccessToken)

		AuthMiddleware()(gc)
		isAuthed := gc.Value(authedCheckCtxKey.name).(bool)

		s.Equal(false, isAuthed)
	})

	s.service.DeleteUserById(context.Background(), targetUser.ID)
}
