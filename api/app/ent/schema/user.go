package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// field.String("id_str"),
		field.String("name").NotEmpty(),
		field.String("screen_name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("profile_image"),
		// field.Int("followers_count"),
		// field.Int("followings_count"),
		// field.Int("likes_count"),
		// field.Int("num_tweets_count"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Tweet.Type),
		edge.To("following", User.Type).From("followers"),
		edge.To("puts", Like.Type),
	}
}
