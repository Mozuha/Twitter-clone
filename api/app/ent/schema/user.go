package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
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
		field.String("name"). // name to be displayed without @ mark
					SchemaType(map[string]string{
				dialect.Postgres: "varchar(50)",
			}).
			NotEmpty().
			Annotations(entgql.OrderField("NAME")),
		field.String("screen_name"). // aka handle name (@screen_name)
						SchemaType(map[string]string{
				dialect.Postgres: "varchar(15)",
			}).
			NotEmpty().
			Unique().
			Annotations(entgql.OrderField("SCREEN_NAME")),
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("password").
			NotEmpty().
			Sensitive(),
		field.String("profile_image").
			Default("images/default.jpg"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entgql.OrderField("UPDATED_AT")),

		// potential fields referring to https://developer.twitter.com/en/docs/twitter-api/v1/data-dictionary/object-model/user
		// field.String("id_str"),
		// field.Int("followers_count"),
		// field.Int("followings_count"),
		// field.Int("likes_count"),
		// field.Int("num_tweets_count"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Tweet.Type),
		edge.To("following", User.Type).From("followers"),
		edge.To("likes", Tweet.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField("users(where: UserWhereInput)"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
