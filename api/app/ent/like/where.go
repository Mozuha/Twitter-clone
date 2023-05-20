// Code generated by ent, DO NOT EDIT.

package like

import (
	"api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUserID, v))
}

// TweetID applies equality check predicate on the "tweet_id" field. It's identical to TweetIDEQ.
func TweetID(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldTweetID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldUserID, v))
}

// TweetIDEQ applies the EQ predicate on the "tweet_id" field.
func TweetIDEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldTweetID, v))
}

// TweetIDNEQ applies the NEQ predicate on the "tweet_id" field.
func TweetIDNEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldTweetID, v))
}

// TweetIDIn applies the In predicate on the "tweet_id" field.
func TweetIDIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldTweetID, vs...))
}

// TweetIDNotIn applies the NotIn predicate on the "tweet_id" field.
func TweetIDNotIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldTweetID, vs...))
}

// TweetIDGT applies the GT predicate on the "tweet_id" field.
func TweetIDGT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldTweetID, v))
}

// TweetIDGTE applies the GTE predicate on the "tweet_id" field.
func TweetIDGTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldTweetID, v))
}

// TweetIDLT applies the LT predicate on the "tweet_id" field.
func TweetIDLT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldTweetID, v))
}

// TweetIDLTE applies the LTE predicate on the "tweet_id" field.
func TweetIDLTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldTweetID, v))
}

// HasPutBy applies the HasEdge predicate on the "put_by" edge.
func HasPutBy() predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PutByTable, PutByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPutByWith applies the HasEdge predicate on the "put_by" edge with a given conditions (other predicates).
func HasPutByWith(preds ...predicate.User) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := newPutByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBelongTo applies the HasEdge predicate on the "belong_to" edge.
func HasBelongTo() predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BelongToTable, BelongToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBelongToWith applies the HasEdge predicate on the "belong_to" edge with a given conditions (other predicates).
func HasBelongToWith(preds ...predicate.Tweet) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := newBelongToStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Like) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		p(s.Not())
	})
}