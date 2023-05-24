// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/like"
	"app/ent/tweet"
	"app/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LikeCreate is the builder for creating a Like entity.
type LikeCreate struct {
	config
	mutation *LikeMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (lc *LikeCreate) SetUserID(i int) *LikeCreate {
	lc.mutation.SetUserID(i)
	return lc
}

// SetTweetID sets the "tweet_id" field.
func (lc *LikeCreate) SetTweetID(i int) *LikeCreate {
	lc.mutation.SetTweetID(i)
	return lc
}

// SetPutByID sets the "put_by" edge to the User entity by ID.
func (lc *LikeCreate) SetPutByID(id int) *LikeCreate {
	lc.mutation.SetPutByID(id)
	return lc
}

// SetPutBy sets the "put_by" edge to the User entity.
func (lc *LikeCreate) SetPutBy(u *User) *LikeCreate {
	return lc.SetPutByID(u.ID)
}

// SetBelongToID sets the "belong_to" edge to the Tweet entity by ID.
func (lc *LikeCreate) SetBelongToID(id int) *LikeCreate {
	lc.mutation.SetBelongToID(id)
	return lc
}

// SetBelongTo sets the "belong_to" edge to the Tweet entity.
func (lc *LikeCreate) SetBelongTo(t *Tweet) *LikeCreate {
	return lc.SetBelongToID(t.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (lc *LikeCreate) Mutation() *LikeMutation {
	return lc.mutation
}

// Save creates the Like in the database.
func (lc *LikeCreate) Save(ctx context.Context) (*Like, error) {
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LikeCreate) SaveX(ctx context.Context) *Like {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LikeCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LikeCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LikeCreate) check() error {
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Like.user_id"`)}
	}
	if _, ok := lc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "Like.tweet_id"`)}
	}
	if _, ok := lc.mutation.PutByID(); !ok {
		return &ValidationError{Name: "put_by", err: errors.New(`ent: missing required edge "Like.put_by"`)}
	}
	if _, ok := lc.mutation.BelongToID(); !ok {
		return &ValidationError{Name: "belong_to", err: errors.New(`ent: missing required edge "Like.belong_to"`)}
	}
	return nil
}

func (lc *LikeCreate) sqlSave(ctx context.Context) (*Like, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LikeCreate) createSpec() (*Like, *sqlgraph.CreateSpec) {
	var (
		_node = &Like{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(like.Table, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.UserID(); ok {
		_spec.SetField(like.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := lc.mutation.TweetID(); ok {
		_spec.SetField(like.FieldTweetID, field.TypeInt, value)
		_node.TweetID = value
	}
	if nodes := lc.mutation.PutByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.PutByTable,
			Columns: []string{like.PutByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_likes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.BelongToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.BelongToTable,
			Columns: []string{like.BelongToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.tweet_has = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LikeCreateBulk is the builder for creating many Like entities in bulk.
type LikeCreateBulk struct {
	config
	builders []*LikeCreate
}

// Save creates the Like entities in the database.
func (lcb *LikeCreateBulk) Save(ctx context.Context) ([]*Like, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Like, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LikeCreateBulk) SaveX(ctx context.Context) []*Like {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LikeCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LikeCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
