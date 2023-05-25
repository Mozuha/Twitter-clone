package app

import (
	"app/db"
	"app/ent"
	"app/ent/user"
	"app/utils"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EntUserServiceTestSuite struct {
	suite.Suite
	db  *ent.Client
	ctx context.Context
}

func (s *EntUserServiceTestSuite) SetupTest() {
	runningEnv, err := utils.LoadEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	entClient, err := db.ConnectTestDB(runningEnv)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	s.db = entClient
	s.ctx = context.Background()
}

func TestEntUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EntUserServiceTestSuite))
}

func (s *EntUserServiceTestSuite) TestGetUsers() {
	users, err := s.db.User.Query().All(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.NotEmpty(users)
}

func (s *EntUserServiceTestSuite) TestGetUserByID() {
	targetUser, err := s.db.User.Query().First(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.Run("success", func() {
		user, err := s.db.User.Get(s.ctx, targetUser.ID)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(targetUser.ID, user.ID)
		s.Equal(targetUser.Email, user.Email)
	})

	s.Run("error/not found", func() {
		_, err := s.db.User.Get(s.ctx, 100)
		if !ent.IsNotFound(err) {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(true, ent.IsNotFound(err))
	})
}

func (s *EntUserServiceTestSuite) TestCreateUser() {
	expectedUser := &ent.User{
		Name:       "test 4",
		ScreenName: "test4",
		Email:      "test4@ymail.ne.jp",
		Password:   "12345",
	}

	s.Run("success", func() {
		user, err := s.db.User.
			Create().
			SetName(expectedUser.Name).
			SetScreenName(expectedUser.ScreenName).
			SetEmail(expectedUser.Email).
			SetPassword(expectedUser.Password).
			Save(s.ctx)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(expectedUser.Name, user.Name)
		s.Equal(expectedUser.ScreenName, user.ScreenName)
		s.Equal(expectedUser.Email, user.Email)
		s.Equal(expectedUser.Password, user.Password)
	})

	s.Run("error/name field (required) is missing", func() {
		err := s.db.User.
			Create().
			SetScreenName(expectedUser.ScreenName).
			SetEmail(expectedUser.Email).
			SetPassword(expectedUser.Password).
			Exec(s.ctx)

		s.Equal(true, s.Error(err))
	})

	s.Run("error/user already exists (email must be unique)", func() {
		err := s.db.User.
			Create().
			SetName(expectedUser.Name).
			SetScreenName(expectedUser.ScreenName).
			SetEmail("test1@gmail.com").
			SetPassword(expectedUser.Password).
			Exec(s.ctx)

		s.Equal(true, s.Error(err))
	})
}

func (s *EntUserServiceTestSuite) TestUpdateUserByID() {
	targetUser, err := s.db.User.Query().Where(user.Email("test2@ymail.ne.jp")).Only(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	expectedUser := &ent.User{
		Name:       "test 2 updated",
		ScreenName: "test2neo",
		Email:      "test2neo@ymail.ne.jp",
		Password:   "newpassword",
	}

	s.Run("success", func() {
		user, err := s.db.User.
			UpdateOneID(targetUser.ID).
			SetName(expectedUser.Name).
			SetScreenName(expectedUser.ScreenName).
			SetEmail(expectedUser.Email).
			SetPassword(expectedUser.Password).
			Save(s.ctx)
		if err != nil {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(expectedUser.Name, user.Name)
		s.Equal(expectedUser.ScreenName, user.ScreenName)
		s.Equal(expectedUser.Email, user.Email)
		s.Equal(expectedUser.Password, user.Password)
	})

	s.Run("error/name field (required) is missing", func() {
		err := s.db.User.
			UpdateOneID(targetUser.ID).
			SetName("").
			SetScreenName(expectedUser.ScreenName).
			SetEmail(expectedUser.Email).
			SetPassword(expectedUser.Password).
			Exec(s.ctx)

		s.Equal(true, s.Error(err))
	})

	s.Run("error/user already exists (email must be unique)", func() {
		err := s.db.User.
			UpdateOneID(targetUser.ID).
			SetName(expectedUser.Name).
			SetScreenName(expectedUser.ScreenName).
			SetEmail("test1@gmail.com").
			SetPassword(expectedUser.Password).
			Exec(s.ctx)

		s.Equal(true, s.Error(err))
	})

	s.Run("error/not found", func() {
		err := s.db.User.
			UpdateOneID(100).
			SetName(expectedUser.Name).
			SetScreenName(expectedUser.ScreenName).
			SetEmail(expectedUser.Email).
			SetPassword(expectedUser.Password).
			Exec(s.ctx)
		if !ent.IsNotFound(err) {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(true, ent.IsNotFound(err))
	})
}

func (s *EntUserServiceTestSuite) TestDeleteUserByID() {
	targetUser, err := s.db.User.Query().Where(user.Email("test3@gmail.com")).Only(s.ctx)
	if err != nil {
		s.Fail("unexpected error occurred: ", err)
	}

	s.Run("success", func() {
		err := s.db.User.DeleteOneID(targetUser.ID).Exec(s.ctx)
		s.NoError(err)
		_, err = s.db.User.Query().Where(user.Email("test3@gmail.com")).Only(s.ctx)
		s.Equal(true, ent.IsNotFound(err))
	})

	s.Run("error/not found", func() {
		err := s.db.User.DeleteOneID(100).Exec(s.ctx)
		s.Equal(true, ent.IsNotFound(err))
	})
}
