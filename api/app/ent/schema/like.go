package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("tweet_id", uuid.UUID{}),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("put_by", User.Type).Unique().Required().Ref("puts"),
		edge.From("belong_to", Tweet.Type).Unique().Required().Ref("has"),
	}
}
