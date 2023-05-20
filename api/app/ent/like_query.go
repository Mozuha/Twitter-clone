// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/like"
	"api/ent/predicate"
	"api/ent/tweet"
	"api/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LikeQuery is the builder for querying Like entities.
type LikeQuery struct {
	config
	ctx          *QueryContext
	order        []like.OrderOption
	inters       []Interceptor
	predicates   []predicate.Like
	withPutBy    *UserQuery
	withBelongTo *TweetQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LikeQuery builder.
func (lq *LikeQuery) Where(ps ...predicate.Like) *LikeQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LikeQuery) Limit(limit int) *LikeQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LikeQuery) Offset(offset int) *LikeQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LikeQuery) Unique(unique bool) *LikeQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LikeQuery) Order(o ...like.OrderOption) *LikeQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryPutBy chains the current query on the "put_by" edge.
func (lq *LikeQuery) QueryPutBy() *UserQuery {
	query := (&UserClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(like.Table, like.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, like.PutByTable, like.PutByColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBelongTo chains the current query on the "belong_to" edge.
func (lq *LikeQuery) QueryBelongTo() *TweetQuery {
	query := (&TweetClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(like.Table, like.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, like.BelongToTable, like.BelongToColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Like entity from the query.
// Returns a *NotFoundError when no Like was found.
func (lq *LikeQuery) First(ctx context.Context) (*Like, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{like.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LikeQuery) FirstX(ctx context.Context) *Like {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Like ID from the query.
// Returns a *NotFoundError when no Like ID was found.
func (lq *LikeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{like.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LikeQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Like entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Like entity is found.
// Returns a *NotFoundError when no Like entities are found.
func (lq *LikeQuery) Only(ctx context.Context) (*Like, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{like.Label}
	default:
		return nil, &NotSingularError{like.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LikeQuery) OnlyX(ctx context.Context) *Like {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Like ID in the query.
// Returns a *NotSingularError when more than one Like ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LikeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{like.Label}
	default:
		err = &NotSingularError{like.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LikeQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Likes.
func (lq *LikeQuery) All(ctx context.Context) ([]*Like, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Like, *LikeQuery]()
	return withInterceptors[[]*Like](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LikeQuery) AllX(ctx context.Context) []*Like {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Like IDs.
func (lq *LikeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(like.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LikeQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LikeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LikeQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LikeQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LikeQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LikeQuery) Clone() *LikeQuery {
	if lq == nil {
		return nil
	}
	return &LikeQuery{
		config:       lq.config,
		ctx:          lq.ctx.Clone(),
		order:        append([]like.OrderOption{}, lq.order...),
		inters:       append([]Interceptor{}, lq.inters...),
		predicates:   append([]predicate.Like{}, lq.predicates...),
		withPutBy:    lq.withPutBy.Clone(),
		withBelongTo: lq.withBelongTo.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithPutBy tells the query-builder to eager-load the nodes that are connected to
// the "put_by" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithPutBy(opts ...func(*UserQuery)) *LikeQuery {
	query := (&UserClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withPutBy = query
	return lq
}

// WithBelongTo tells the query-builder to eager-load the nodes that are connected to
// the "belong_to" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithBelongTo(opts ...func(*TweetQuery)) *LikeQuery {
	query := (&TweetClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withBelongTo = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Like.Query().
//		GroupBy(like.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LikeQuery) GroupBy(field string, fields ...string) *LikeGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LikeGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = like.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.Like.Query().
//		Select(like.FieldUserID).
//		Scan(ctx, &v)
func (lq *LikeQuery) Select(fields ...string) *LikeSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LikeSelect{LikeQuery: lq}
	sbuild.label = like.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LikeSelect configured with the given aggregations.
func (lq *LikeQuery) Aggregate(fns ...AggregateFunc) *LikeSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !like.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Like, error) {
	var (
		nodes       = []*Like{}
		withFKs     = lq.withFKs
		_spec       = lq.querySpec()
		loadedTypes = [2]bool{
			lq.withPutBy != nil,
			lq.withBelongTo != nil,
		}
	)
	if lq.withPutBy != nil || lq.withBelongTo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, like.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Like).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Like{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withPutBy; query != nil {
		if err := lq.loadPutBy(ctx, query, nodes, nil,
			func(n *Like, e *User) { n.Edges.PutBy = e }); err != nil {
			return nil, err
		}
	}
	if query := lq.withBelongTo; query != nil {
		if err := lq.loadBelongTo(ctx, query, nodes, nil,
			func(n *Like, e *Tweet) { n.Edges.BelongTo = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LikeQuery) loadPutBy(ctx context.Context, query *UserQuery, nodes []*Like, init func(*Like), assign func(*Like, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Like)
	for i := range nodes {
		if nodes[i].user_puts == nil {
			continue
		}
		fk := *nodes[i].user_puts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_puts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lq *LikeQuery) loadBelongTo(ctx context.Context, query *TweetQuery, nodes []*Like, init func(*Like), assign func(*Like, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Like)
	for i := range nodes {
		if nodes[i].tweet_has == nil {
			continue
		}
		fk := *nodes[i].tweet_has
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tweet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tweet_has" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lq *LikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, like.FieldID)
		for i := range fields {
			if fields[i] != like.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(like.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = like.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LikeGroupBy is the group-by builder for Like entities.
type LikeGroupBy struct {
	selector
	build *LikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LikeGroupBy) Aggregate(fns ...AggregateFunc) *LikeGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikeQuery, *LikeGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LikeGroupBy) sqlScan(ctx context.Context, root *LikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LikeSelect is the builder for selecting fields of Like entities.
type LikeSelect struct {
	*LikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LikeSelect) Aggregate(fns ...AggregateFunc) *LikeSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LikeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikeQuery, *LikeSelect](ctx, ls.LikeQuery, ls, ls.inters, v)
}

func (ls *LikeSelect) sqlScan(ctx context.Context, root *LikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
