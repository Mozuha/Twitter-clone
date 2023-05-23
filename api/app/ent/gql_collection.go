// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/like"
	"app/ent/tweet"
	"app/ent/user"
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *LikeQuery) CollectFields(ctx context.Context, satisfies ...string) (*LikeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return l, nil
	}
	if err := l.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *LikeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(like.Columns))
		selectedFields = []string{like.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "putBy":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: l.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			l.withPutBy = query
		case "belongTo":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TweetClient{config: l.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, tweetImplementors)...); err != nil {
				return err
			}
			l.withBelongTo = query
		case "userID":
			if _, ok := fieldSeen[like.FieldUserID]; !ok {
				selectedFields = append(selectedFields, like.FieldUserID)
				fieldSeen[like.FieldUserID] = struct{}{}
			}
		case "tweetID":
			if _, ok := fieldSeen[like.FieldTweetID]; !ok {
				selectedFields = append(selectedFields, like.FieldTweetID)
				fieldSeen[like.FieldTweetID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		l.Select(selectedFields...)
	}
	return nil
}

type likePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LikePaginateOption
}

func newLikePaginateArgs(rv map[string]any) *likePaginateArgs {
	args := &likePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TweetQuery) CollectFields(ctx context.Context, satisfies ...string) (*TweetQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TweetQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(tweet.Columns))
		selectedFields = []string{tweet.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "postedBy":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			t.withPostedBy = query
		case "child":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TweetClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, tweetImplementors)...); err != nil {
				return err
			}
			t.WithNamedChild(alias, func(wq *TweetQuery) {
				*wq = *query
			})
		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TweetClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, tweetImplementors)...); err != nil {
				return err
			}
			t.WithNamedParent(alias, func(wq *TweetQuery) {
				*wq = *query
			})
		case "has":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&LikeClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, likeImplementors)...); err != nil {
				return err
			}
			t.WithNamedHas(alias, func(wq *LikeQuery) {
				*wq = *query
			})
		case "text":
			if _, ok := fieldSeen[tweet.FieldText]; !ok {
				selectedFields = append(selectedFields, tweet.FieldText)
				fieldSeen[tweet.FieldText] = struct{}{}
			}
		case "parentID":
			if _, ok := fieldSeen[tweet.FieldParentID]; !ok {
				selectedFields = append(selectedFields, tweet.FieldParentID)
				fieldSeen[tweet.FieldParentID] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[tweet.FieldUserID]; !ok {
				selectedFields = append(selectedFields, tweet.FieldUserID)
				fieldSeen[tweet.FieldUserID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[tweet.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, tweet.FieldCreatedAt)
				fieldSeen[tweet.FieldCreatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type tweetPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TweetPaginateOption
}

func newTweetPaginateArgs(rv map[string]any) *tweetPaginateArgs {
	args := &tweetPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "posts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TweetClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, tweetImplementors)...); err != nil {
				return err
			}
			u.WithNamedPosts(alias, func(wq *TweetQuery) {
				*wq = *query
			})
		case "followers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			u.WithNamedFollowers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "following":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			u.WithNamedFollowing(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "puts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&LikeClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, likeImplementors)...); err != nil {
				return err
			}
			u.WithNamedPuts(alias, func(wq *LikeQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "screenName":
			if _, ok := fieldSeen[user.FieldScreenName]; !ok {
				selectedFields = append(selectedFields, user.FieldScreenName)
				fieldSeen[user.FieldScreenName] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "password":
			if _, ok := fieldSeen[user.FieldPassword]; !ok {
				selectedFields = append(selectedFields, user.FieldPassword)
				fieldSeen[user.FieldPassword] = struct{}{}
			}
		case "profileImage":
			if _, ok := fieldSeen[user.FieldProfileImage]; !ok {
				selectedFields = append(selectedFields, user.FieldProfileImage)
				fieldSeen[user.FieldProfileImage] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
