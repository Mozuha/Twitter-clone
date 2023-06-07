package utils

import (
	"app/middlewares"
	"context"
	"errors"
)

func CheckIsAuthedFromCtx(ctx context.Context) error {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return err
	}

	if isAuthed := middlewares.ForContext(gc); !isAuthed {
		return errors.New("request not authorized")
	} else {
		return nil
	}
}
