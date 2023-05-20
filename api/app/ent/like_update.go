// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/like"
	"api/ent/predicate"
	"api/ent/tweet"
	"api/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LikeUpdate is the builder for updating Like entities.
type LikeUpdate struct {
	config
	hooks    []Hook
	mutation *LikeMutation
}

// Where appends a list predicates to the LikeUpdate builder.
func (lu *LikeUpdate) Where(ps ...predicate.Like) *LikeUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUserID sets the "user_id" field.
func (lu *LikeUpdate) SetUserID(u uuid.UUID) *LikeUpdate {
	lu.mutation.SetUserID(u)
	return lu
}

// SetTweetID sets the "tweet_id" field.
func (lu *LikeUpdate) SetTweetID(u uuid.UUID) *LikeUpdate {
	lu.mutation.SetTweetID(u)
	return lu
}

// SetPutByID sets the "put_by" edge to the User entity by ID.
func (lu *LikeUpdate) SetPutByID(id int) *LikeUpdate {
	lu.mutation.SetPutByID(id)
	return lu
}

// SetPutBy sets the "put_by" edge to the User entity.
func (lu *LikeUpdate) SetPutBy(u *User) *LikeUpdate {
	return lu.SetPutByID(u.ID)
}

// SetBelongToID sets the "belong_to" edge to the Tweet entity by ID.
func (lu *LikeUpdate) SetBelongToID(id int) *LikeUpdate {
	lu.mutation.SetBelongToID(id)
	return lu
}

// SetBelongTo sets the "belong_to" edge to the Tweet entity.
func (lu *LikeUpdate) SetBelongTo(t *Tweet) *LikeUpdate {
	return lu.SetBelongToID(t.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (lu *LikeUpdate) Mutation() *LikeMutation {
	return lu.mutation
}

// ClearPutBy clears the "put_by" edge to the User entity.
func (lu *LikeUpdate) ClearPutBy() *LikeUpdate {
	lu.mutation.ClearPutBy()
	return lu
}

// ClearBelongTo clears the "belong_to" edge to the Tweet entity.
func (lu *LikeUpdate) ClearBelongTo() *LikeUpdate {
	lu.mutation.ClearBelongTo()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LikeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LikeUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LikeUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LikeUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LikeUpdate) check() error {
	if _, ok := lu.mutation.PutByID(); lu.mutation.PutByCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.put_by"`)
	}
	if _, ok := lu.mutation.BelongToID(); lu.mutation.BelongToCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.belong_to"`)
	}
	return nil
}

func (lu *LikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UserID(); ok {
		_spec.SetField(like.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := lu.mutation.TweetID(); ok {
		_spec.SetField(like.FieldTweetID, field.TypeUUID, value)
	}
	if lu.mutation.PutByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.PutByIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.BelongToCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.BelongToIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LikeUpdateOne is the builder for updating a single Like entity.
type LikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LikeMutation
}

// SetUserID sets the "user_id" field.
func (luo *LikeUpdateOne) SetUserID(u uuid.UUID) *LikeUpdateOne {
	luo.mutation.SetUserID(u)
	return luo
}

// SetTweetID sets the "tweet_id" field.
func (luo *LikeUpdateOne) SetTweetID(u uuid.UUID) *LikeUpdateOne {
	luo.mutation.SetTweetID(u)
	return luo
}

// SetPutByID sets the "put_by" edge to the User entity by ID.
func (luo *LikeUpdateOne) SetPutByID(id int) *LikeUpdateOne {
	luo.mutation.SetPutByID(id)
	return luo
}

// SetPutBy sets the "put_by" edge to the User entity.
func (luo *LikeUpdateOne) SetPutBy(u *User) *LikeUpdateOne {
	return luo.SetPutByID(u.ID)
}

// SetBelongToID sets the "belong_to" edge to the Tweet entity by ID.
func (luo *LikeUpdateOne) SetBelongToID(id int) *LikeUpdateOne {
	luo.mutation.SetBelongToID(id)
	return luo
}

// SetBelongTo sets the "belong_to" edge to the Tweet entity.
func (luo *LikeUpdateOne) SetBelongTo(t *Tweet) *LikeUpdateOne {
	return luo.SetBelongToID(t.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (luo *LikeUpdateOne) Mutation() *LikeMutation {
	return luo.mutation
}

// ClearPutBy clears the "put_by" edge to the User entity.
func (luo *LikeUpdateOne) ClearPutBy() *LikeUpdateOne {
	luo.mutation.ClearPutBy()
	return luo
}

// ClearBelongTo clears the "belong_to" edge to the Tweet entity.
func (luo *LikeUpdateOne) ClearBelongTo() *LikeUpdateOne {
	luo.mutation.ClearBelongTo()
	return luo
}

// Where appends a list predicates to the LikeUpdate builder.
func (luo *LikeUpdateOne) Where(ps ...predicate.Like) *LikeUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LikeUpdateOne) Select(field string, fields ...string) *LikeUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Like entity.
func (luo *LikeUpdateOne) Save(ctx context.Context) (*Like, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LikeUpdateOne) SaveX(ctx context.Context) *Like {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LikeUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LikeUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LikeUpdateOne) check() error {
	if _, ok := luo.mutation.PutByID(); luo.mutation.PutByCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.put_by"`)
	}
	if _, ok := luo.mutation.BelongToID(); luo.mutation.BelongToCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.belong_to"`)
	}
	return nil
}

func (luo *LikeUpdateOne) sqlSave(ctx context.Context) (_node *Like, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Like.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, like.FieldID)
		for _, f := range fields {
			if !like.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != like.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UserID(); ok {
		_spec.SetField(like.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := luo.mutation.TweetID(); ok {
		_spec.SetField(like.FieldTweetID, field.TypeUUID, value)
	}
	if luo.mutation.PutByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.PutByIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.BelongToCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.BelongToIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Like{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}