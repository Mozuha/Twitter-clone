// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/schema"
	"app/ent/tweet"
	"app/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tweetFields := schema.Tweet{}.Fields()
	_ = tweetFields
	// tweetDescText is the schema descriptor for text field.
	tweetDescText := tweetFields[0].Descriptor()
	// tweet.TextValidator is a validator for the "text" field. It is called by the builders before save.
	tweet.TextValidator = tweetDescText.Validators[0].(func(string) error)
	// tweetDescCreatedAt is the schema descriptor for created_at field.
	tweetDescCreatedAt := tweetFields[3].Descriptor()
	// tweet.DefaultCreatedAt holds the default value on creation for the created_at field.
	tweet.DefaultCreatedAt = tweetDescCreatedAt.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescScreenName is the schema descriptor for screen_name field.
	userDescScreenName := userFields[1].Descriptor()
	// user.ScreenNameValidator is a validator for the "screen_name" field. It is called by the builders before save.
	user.ScreenNameValidator = userDescScreenName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescProfileImage is the schema descriptor for profile_image field.
	userDescProfileImage := userFields[4].Descriptor()
	// user.DefaultProfileImage holds the default value on creation for the profile_image field.
	user.DefaultProfileImage = userDescProfileImage.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
