package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		field.UUID("parent_id", uuid.UUID{}).Nillable().Optional(),
		field.UUID("user_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posted_by", User.Type).Unique().Required().Ref("posts"),
		edge.To("parent", Tweet.Type).From("child"),
		edge.To("has", Like.Type),
	}
}
