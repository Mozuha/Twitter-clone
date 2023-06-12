package services

import (
	"app/db"
	"app/ent"
	"app/ent/user"
	"app/utils"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

const USER_NOT_FOUND_ERROR = "input: ent: user not found"

type UserServiceTestSuite struct {
	suite.Suite
	db      *ent.Client
	ctx     context.Context
	service Services
}

func (s *UserServiceTestSuite) SetupTest() {
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

func (s *UserServiceTestSuite) TearDownTest() {
	s.db.Close()
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (s *UserServiceTestSuite) TestGetUsers() {
	users, err := s.service.GetUsers(s.ctx, &ent.UserWhereInput{})

	s.NotEmpty(users)
	s.NoError(err)
}

func (s *UserServiceTestSuite) TestGetUserByID() {
	targetUser, err := s.db.User.Query().First(s.ctx)
	if err != nil {
		s.Fail("failed to get user to be used: ", err)
	}

	s.Run("success", func() {
		user, err := s.service.GetUsers(s.ctx, &ent.UserWhereInput{ID: &targetUser.ID})

		s.Equal(targetUser.ID, user[0].ID)
		s.Equal(targetUser.Email, user[0].Email)
		s.NoError(err)
	})

	s.Run("error/not found", func() {
		notExistingId := 100
		_, err := s.service.GetUsers(s.ctx, &ent.UserWhereInput{ID: &notExistingId})
		if err.Error() != USER_NOT_FOUND_ERROR {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(true, err.Error() == USER_NOT_FOUND_ERROR)
	})
}

func (s *UserServiceTestSuite) TestCreateUser() {
	expectedUser := &ent.User{
		Name:       "test 5",
		ScreenName: "test5",
		Email:      "test5@ymail.ne.jp",
		Password:   "12345",
	}

	s.Run("success", func() {
		input := ent.CreateUserInput{
			Name:       expectedUser.Name,
			ScreenName: expectedUser.ScreenName,
			Email:      expectedUser.Email,
			Password:   expectedUser.Password,
		}

		user, err := s.service.CreateUser(s.ctx, input)
		pwIntegrityErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(expectedUser.Password))

		s.Equal(expectedUser.Name, user.Name)
		s.Equal(expectedUser.ScreenName, user.ScreenName)
		s.Equal(expectedUser.Email, user.Email)
		s.NoError(err)
		s.NoError(pwIntegrityErr)

		s.service.DeleteUserById(s.ctx, user.ID)
	})

	s.Run("error/name field (required) is missing", func() {
		input := ent.CreateUserInput{
			ScreenName: expectedUser.ScreenName,
			Email:      expectedUser.Email,
			Password:   expectedUser.Password,
		}

		_, err := s.service.CreateUser(s.ctx, input)

		s.Error(err)
	})

	s.Run("error/user already exists (email must be unique)", func() {
		input := ent.CreateUserInput{
			Name:       expectedUser.Name,
			ScreenName: expectedUser.ScreenName,
			Email:      "test1@gmail.com",
			Password:   expectedUser.Password,
		}

		_, err := s.service.CreateUser(s.ctx, input)

		s.Error(err)
	})

	// TODO: Name, ScreenName length check
}

func (s *UserServiceTestSuite) TestUpdateUserByID() {
	targetUser, err := s.db.User.Query().Where(user.Email("test2@ymail.ne.jp")).Only(s.ctx)
	if err != nil {
		s.Fail("failed to get user to be updated: ", err)
	}

	expectedUser := &ent.User{
		Name:       "test 2 updated",
		ScreenName: "test2neo",
		Email:      "test2neo@ymail.ne.jp",
		Password:   "newpassword",
	}

	s.Run("success", func() {
		input := ent.UpdateUserInput{
			Name:       &expectedUser.Name,
			ScreenName: &expectedUser.ScreenName,
			Email:      &expectedUser.Email,
			Password:   &expectedUser.Password,
		}

		user, err := s.service.UpdateUserById(s.ctx, targetUser.ID, input)
		pwIntegrityErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(expectedUser.Password))

		s.Equal(expectedUser.Name, user.Name)
		s.Equal(expectedUser.ScreenName, user.ScreenName)
		s.Equal(expectedUser.Email, user.Email)
		s.NoError(err)
		s.NoError(pwIntegrityErr)
	})

	s.Run("error/user already exists (email must be unique)", func() {
		existingEmail := "test1@gmail.com"
		input := ent.UpdateUserInput{
			Name:       &expectedUser.Name,
			ScreenName: &expectedUser.ScreenName,
			Email:      &existingEmail,
			Password:   &expectedUser.Password,
		}

		_, err := s.service.UpdateUserById(s.ctx, targetUser.ID, input)

		s.Error(err)
	})

	s.Run("error/not found", func() {
		input := ent.UpdateUserInput{
			Name:       &expectedUser.Name,
			ScreenName: &expectedUser.ScreenName,
			Email:      &expectedUser.Email,
			Password:   &expectedUser.Password,
		}

		_, err := s.service.UpdateUserById(s.ctx, 100, input)

		if err.Error() != USER_NOT_FOUND_ERROR {
			s.Fail("unexpected error occurred: ", err)
		}

		s.Equal(true, err.Error() == USER_NOT_FOUND_ERROR)
	})
}

func (s *UserServiceTestSuite) TestDeleteUserByID() {
	targetUser := &ent.User{
		Name:       "test 4",
		ScreenName: "test4",
		Email:      "test4@ymail.ne.jp",
		Password:   "tobedeleted",
	}

	input := ent.CreateUserInput{
		Name:       targetUser.Name,
		ScreenName: targetUser.ScreenName,
		Email:      targetUser.Email,
		Password:   targetUser.Password,
	}

	user, err := s.service.CreateUser(s.ctx, input)
	if err != nil {
		s.Fail("failed to create user to be deleted: ", err)
	}

	s.Run("success", func() {
		isDeleted, err := s.service.DeleteUserById(s.ctx, user.ID)
		s.Equal(true, *isDeleted)
		s.NoError(err)

		_, err = s.service.GetUsers(s.ctx, &ent.UserWhereInput{Email: &targetUser.Email})
		s.Equal(true, err.Error() == USER_NOT_FOUND_ERROR)
	})

	s.Run("error/not found", func() {
		_, err := s.service.DeleteUserById(s.ctx, 100)
		s.Equal(true, err.Error() == USER_NOT_FOUND_ERROR)
	})
}

func (s *UserServiceTestSuite) TestCheckEmailExists() {
	s.Run("success", func() {
		isExist, _ := s.service.CheckEmailExists(s.ctx, "test1@gmail.com")
		s.Equal(true, *isExist)

		isExist, _ = s.service.CheckEmailExists(s.ctx, "not.existing@gmail.com")
		s.Equal(false, *isExist)
	})
}

func (s *UserServiceTestSuite) TestCheckScreenNameExists() {
	s.Run("success", func() {
		isExist, _ := s.service.CheckScreenNameExists(s.ctx, "test1")
		s.Equal(true, *isExist)

		isExist, _ = s.service.CheckScreenNameExists(s.ctx, "notexisting")
		s.Equal(false, *isExist)
	})
}

// TODO: user posts tweet? <- create tweet func is tested in tweet service test
// TODO: fetch user's tweets? <- get all tweets regardless of user func is tested  in tweet service test
// TODO: user likes tweet
// TODO: fetch user's likes
// TODO: user follows user
// TODO: fetch user's followings/followers
