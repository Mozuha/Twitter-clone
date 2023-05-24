package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
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
		field.String("text").NotEmpty(),
		field.Int("parent_id").Nillable().Optional(),
		field.Int("user_id"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posted_by", User.Type).Unique().Required().Ref("tweets"),
		edge.To("parent", Tweet.Type).From("child"),
		edge.To("has", Like.Type),
	}
}

func (Tweet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
