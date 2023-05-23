// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (l *Like) PutBy(ctx context.Context) (*User, error) {
	result, err := l.Edges.PutByOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryPutBy().Only(ctx)
	}
	return result, err
}

func (l *Like) BelongTo(ctx context.Context) (*Tweet, error) {
	result, err := l.Edges.BelongToOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryBelongTo().Only(ctx)
	}
	return result, err
}

func (t *Tweet) PostedBy(ctx context.Context) (*User, error) {
	result, err := t.Edges.PostedByOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryPostedBy().Only(ctx)
	}
	return result, err
}

func (t *Tweet) Child(ctx context.Context) (result []*Tweet, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedChild(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ChildOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryChild().All(ctx)
	}
	return result, err
}

func (t *Tweet) Parent(ctx context.Context) (result []*Tweet, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedParent(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ParentOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryParent().All(ctx)
	}
	return result, err
}

func (t *Tweet) Has(ctx context.Context) (result []*Like, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedHas(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.HasOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryHas().All(ctx)
	}
	return result, err
}

func (u *User) Posts(ctx context.Context) (result []*Tweet, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedPosts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.PostsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryPosts().All(ctx)
	}
	return result, err
}

func (u *User) Followers(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedFollowers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.FollowersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryFollowers().All(ctx)
	}
	return result, err
}

func (u *User) Following(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedFollowing(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.FollowingOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryFollowing().All(ctx)
	}
	return result, err
}

func (u *User) Puts(ctx context.Context) (result []*Like, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedPuts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.PutsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryPuts().All(ctx)
	}
	return result, err
}