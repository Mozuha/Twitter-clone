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

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		// field.String("id_str"),
		field.String("text").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(50)",
			}).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now()).
			Immutable().
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posted_by", User.Type).
			Unique().
			Required().
			Ref("posts"),
		// one parent many child, one child one parent
		edge.To("parent", Tweet.Type).
			Unique().
			From("children"),
		edge.From("liked_by", User.Type).
			Ref("likes"),
	}
}

func (Tweet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField("tweets(where: TweetWhereInput)"),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
