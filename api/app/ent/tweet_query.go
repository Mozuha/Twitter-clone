// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/like"
	"app/ent/predicate"
	"app/ent/tweet"
	"app/ent/user"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TweetQuery is the builder for querying Tweet entities.
type TweetQuery struct {
	config
	ctx             *QueryContext
	order           []tweet.OrderOption
	inters          []Interceptor
	predicates      []predicate.Tweet
	withPostedBy    *UserQuery
	withChild       *TweetQuery
	withParent      *TweetQuery
	withHas         *LikeQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*Tweet) error
	withNamedChild  map[string]*TweetQuery
	withNamedParent map[string]*TweetQuery
	withNamedHas    map[string]*LikeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TweetQuery builder.
func (tq *TweetQuery) Where(ps ...predicate.Tweet) *TweetQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TweetQuery) Limit(limit int) *TweetQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TweetQuery) Offset(offset int) *TweetQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TweetQuery) Unique(unique bool) *TweetQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TweetQuery) Order(o ...tweet.OrderOption) *TweetQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryPostedBy chains the current query on the "posted_by" edge.
func (tq *TweetQuery) QueryPostedBy() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tweet.PostedByTable, tweet.PostedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChild chains the current query on the "child" edge.
func (tq *TweetQuery) QueryChild() *TweetQuery {
	query := (&TweetClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.ChildTable, tweet.ChildPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (tq *TweetQuery) QueryParent() *TweetQuery {
	query := (&TweetClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, tweet.ParentTable, tweet.ParentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHas chains the current query on the "has" edge.
func (tq *TweetQuery) QueryHas() *LikeQuery {
	query := (&LikeClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(like.Table, like.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tweet.HasTable, tweet.HasColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tweet entity from the query.
// Returns a *NotFoundError when no Tweet was found.
func (tq *TweetQuery) First(ctx context.Context) (*Tweet, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tweet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TweetQuery) FirstX(ctx context.Context) *Tweet {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tweet ID from the query.
// Returns a *NotFoundError when no Tweet ID was found.
func (tq *TweetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tweet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TweetQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tweet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tweet entity is found.
// Returns a *NotFoundError when no Tweet entities are found.
func (tq *TweetQuery) Only(ctx context.Context) (*Tweet, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tweet.Label}
	default:
		return nil, &NotSingularError{tweet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TweetQuery) OnlyX(ctx context.Context) *Tweet {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tweet ID in the query.
// Returns a *NotSingularError when more than one Tweet ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TweetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tweet.Label}
	default:
		err = &NotSingularError{tweet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TweetQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tweets.
func (tq *TweetQuery) All(ctx context.Context) ([]*Tweet, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tweet, *TweetQuery]()
	return withInterceptors[[]*Tweet](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TweetQuery) AllX(ctx context.Context) []*Tweet {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tweet IDs.
func (tq *TweetQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(tweet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TweetQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TweetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TweetQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TweetQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TweetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TweetQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TweetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TweetQuery) Clone() *TweetQuery {
	if tq == nil {
		return nil
	}
	return &TweetQuery{
		config:       tq.config,
		ctx:          tq.ctx.Clone(),
		order:        append([]tweet.OrderOption{}, tq.order...),
		inters:       append([]Interceptor{}, tq.inters...),
		predicates:   append([]predicate.Tweet{}, tq.predicates...),
		withPostedBy: tq.withPostedBy.Clone(),
		withChild:    tq.withChild.Clone(),
		withParent:   tq.withParent.Clone(),
		withHas:      tq.withHas.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithPostedBy tells the query-builder to eager-load the nodes that are connected to
// the "posted_by" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithPostedBy(opts ...func(*UserQuery)) *TweetQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withPostedBy = query
	return tq
}

// WithChild tells the query-builder to eager-load the nodes that are connected to
// the "child" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithChild(opts ...func(*TweetQuery)) *TweetQuery {
	query := (&TweetClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withChild = query
	return tq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithParent(opts ...func(*TweetQuery)) *TweetQuery {
	query := (&TweetClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withParent = query
	return tq
}

// WithHas tells the query-builder to eager-load the nodes that are connected to
// the "has" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithHas(opts ...func(*LikeQuery)) *TweetQuery {
	query := (&LikeClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withHas = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Tweet.Query().
//		GroupBy(tweet.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TweetQuery) GroupBy(field string, fields ...string) *TweetGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TweetGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = tweet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Tweet.Query().
//		Select(tweet.FieldText).
//		Scan(ctx, &v)
func (tq *TweetQuery) Select(fields ...string) *TweetSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TweetSelect{TweetQuery: tq}
	sbuild.label = tweet.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TweetSelect configured with the given aggregations.
func (tq *TweetQuery) Aggregate(fns ...AggregateFunc) *TweetSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TweetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !tweet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TweetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tweet, error) {
	var (
		nodes       = []*Tweet{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [4]bool{
			tq.withPostedBy != nil,
			tq.withChild != nil,
			tq.withParent != nil,
			tq.withHas != nil,
		}
	)
	if tq.withPostedBy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tweet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tweet{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withPostedBy; query != nil {
		if err := tq.loadPostedBy(ctx, query, nodes, nil,
			func(n *Tweet, e *User) { n.Edges.PostedBy = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withChild; query != nil {
		if err := tq.loadChild(ctx, query, nodes,
			func(n *Tweet) { n.Edges.Child = []*Tweet{} },
			func(n *Tweet, e *Tweet) { n.Edges.Child = append(n.Edges.Child, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withParent; query != nil {
		if err := tq.loadParent(ctx, query, nodes,
			func(n *Tweet) { n.Edges.Parent = []*Tweet{} },
			func(n *Tweet, e *Tweet) { n.Edges.Parent = append(n.Edges.Parent, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withHas; query != nil {
		if err := tq.loadHas(ctx, query, nodes,
			func(n *Tweet) { n.Edges.Has = []*Like{} },
			func(n *Tweet, e *Like) { n.Edges.Has = append(n.Edges.Has, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedChild {
		if err := tq.loadChild(ctx, query, nodes,
			func(n *Tweet) { n.appendNamedChild(name) },
			func(n *Tweet, e *Tweet) { n.appendNamedChild(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedParent {
		if err := tq.loadParent(ctx, query, nodes,
			func(n *Tweet) { n.appendNamedParent(name) },
			func(n *Tweet, e *Tweet) { n.appendNamedParent(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedHas {
		if err := tq.loadHas(ctx, query, nodes,
			func(n *Tweet) { n.appendNamedHas(name) },
			func(n *Tweet, e *Like) { n.appendNamedHas(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TweetQuery) loadPostedBy(ctx context.Context, query *UserQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Tweet)
	for i := range nodes {
		if nodes[i].user_posts == nil {
			continue
		}
		fk := *nodes[i].user_posts
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
			return fmt.Errorf(`unexpected foreign-key "user_posts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TweetQuery) loadChild(ctx context.Context, query *TweetQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *Tweet)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Tweet)
	nids := make(map[int]map[*Tweet]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tweet.ChildTable)
		s.Join(joinT).On(s.C(tweet.FieldID), joinT.C(tweet.ChildPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tweet.ChildPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tweet.ChildPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tweet]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tweet](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "child" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (tq *TweetQuery) loadParent(ctx context.Context, query *TweetQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *Tweet)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Tweet)
	nids := make(map[int]map[*Tweet]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tweet.ParentTable)
		s.Join(joinT).On(s.C(tweet.FieldID), joinT.C(tweet.ParentPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(tweet.ParentPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tweet.ParentPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tweet]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tweet](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "parent" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (tq *TweetQuery) loadHas(ctx context.Context, query *LikeQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *Like)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Tweet)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Like(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tweet.HasColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.tweet_has
		if fk == nil {
			return fmt.Errorf(`foreign-key "tweet_has" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "tweet_has" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TweetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TweetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for i := range fields {
			if fields[i] != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TweetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tweet.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = tweet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedChild tells the query-builder to eager-load the nodes that are connected to the "child"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithNamedChild(name string, opts ...func(*TweetQuery)) *TweetQuery {
	query := (&TweetClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedChild == nil {
		tq.withNamedChild = make(map[string]*TweetQuery)
	}
	tq.withNamedChild[name] = query
	return tq
}

// WithNamedParent tells the query-builder to eager-load the nodes that are connected to the "parent"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithNamedParent(name string, opts ...func(*TweetQuery)) *TweetQuery {
	query := (&TweetClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedParent == nil {
		tq.withNamedParent = make(map[string]*TweetQuery)
	}
	tq.withNamedParent[name] = query
	return tq
}

// WithNamedHas tells the query-builder to eager-load the nodes that are connected to the "has"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TweetQuery) WithNamedHas(name string, opts ...func(*LikeQuery)) *TweetQuery {
	query := (&LikeClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedHas == nil {
		tq.withNamedHas = make(map[string]*LikeQuery)
	}
	tq.withNamedHas[name] = query
	return tq
}

// TweetGroupBy is the group-by builder for Tweet entities.
type TweetGroupBy struct {
	selector
	build *TweetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TweetGroupBy) Aggregate(fns ...AggregateFunc) *TweetGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TweetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetQuery, *TweetGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TweetGroupBy) sqlScan(ctx context.Context, root *TweetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TweetSelect is the builder for selecting fields of Tweet entities.
type TweetSelect struct {
	*TweetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TweetSelect) Aggregate(fns ...AggregateFunc) *TweetSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TweetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetQuery, *TweetSelect](ctx, ts.TweetQuery, ts, ts.inters, v)
}

func (ts *TweetSelect) sqlScan(ctx context.Context, root *TweetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
