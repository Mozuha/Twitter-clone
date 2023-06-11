package utils

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Status string

const (
	NOT_FOUND             = Status("NOT_FOUND")
	UNAUTHORIZED          = Status("UNAUTHORIZED")
	INTERNAL_SERVER_ERROR = Status("INTERNAL_SERVER_ERROR")
)

func CreateGqlErr(ctx context.Context, err error, code Status, userMsg string) *gqlerror.Error {
	gErr := &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code": code,
		},
	}

	if userMsg != "" {
		gErr.Extensions["userMessage"] = userMsg
	}

	return gErr
}
