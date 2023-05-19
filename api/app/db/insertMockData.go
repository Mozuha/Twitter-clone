package db

import (
	"api/ent"
	"api/models"
	"context"
	"fmt"
	"time"
)

func InsertMockData(ctx context.Context, client *ent.Client) error {
	usersMock := []models.User{
		{Id: 1, ScreenName: "test1", Username: "test 1", Email: "test1@gmail.com", Password: "pass", ProfileImage: "images/test1.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Id: 2, ScreenName: "test2", Username: "test 2", Email: "test2@ymail.ne.jp", Password: "word", ProfileImage: "images/test2.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Id: 3, ScreenName: "test3", Username: "test 3", Email: "test3@gmail.com", Password: "password", ProfileImage: "images/test3.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	bulk := make([]*ent.UserCreate, len(usersMock))
	for i, user := range usersMock {
		bulk[i] = client.User.
			Create().
			SetName(user.Username).
			SetScreenName(user.ScreenName).
			SetEmail(user.Email).
			SetPassword(user.Password).
			SetProfileImage(user.ProfileImage).
			SetCreatedAt(user.CreatedAt).
			SetUpdatedAt(user.UpdatedAt)
	}

	// Insert mock users if not exist; otherwise do nothing
	err := client.User.CreateBulk(bulk...).OnConflictColumns("email").DoNothing().Exec(ctx)
	if err != nil {
		return fmt.Errorf("creating mock users: %w", err)
	}
	return nil
}
