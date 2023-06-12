package middlewares

import (
	"app/auth"
	"app/utils"
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
// https://gqlgen.com/recipes/authentication/
var authedCheckCtxKey = &contextKey{"isAuthed"}

type contextKey struct {
	name string
}

// Rather than shutting down the request at this point, this middleware will just set the isAuthed value and let the resolvers decide what to do with it
// so that each resolver can be either public or private, just like the endpoints in a REST API
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// add gin.Context to context.Context so the resolvers can access the gin.Context in order to retrieve the isAuthed value
		// https://gqlgen.com/recipes/gin/
		ctx.Set(authedCheckCtxKey.name, false)
		c := context.WithValue(ctx.Request.Context(), "GinContextKey", ctx)
		ctx.Request = ctx.Request.WithContext(c)

		// Bearer token will be shown like `Authorization: Bearer <token>` in http header
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Error(errors.New("Authorization header is required"))
			ctx.Next()
			return
		}

		// token validation does not distinguish between access and refresh tokens
		// judging whether the token expiration message is for access or refresh token is done by the frontend by checking which token it used to make the request
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.Error(fmt.Errorf("Invalid token: %w", err))
			ctx.Next()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		uId := claims["user_id"].(string)
		sId := claims["session_id"].(string)

		// even if the token is valid, the user cannot be authenticated if there is no associated session
		// this prevents the user from being authenticated with valid token after logged out
		session := sessions.Default(ctx)
		sessionId := session.ID()
		userId := session.Get("user")
		if userId == nil || sId != sessionId || uId != userId {
			ctx.Next()
			return
		}

		// user exists (as session value which was set upon successful signin exists), token verified
		// thus this request is authenticated; overwrite context value
		ctx.Set(authedCheckCtxKey.name, true)
		c = context.WithValue(ctx.Request.Context(), "GinContextKey", ctx)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}

// Check the isAuthed value from the context. REQUIRES Middleware to have run.
func CheckIsAuthedFromCtx(ctx context.Context) error {
	gc, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return err
	}

	if isAuthed := gc.Value(authedCheckCtxKey.name).(bool); !isAuthed {
		gErr := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "User not authenticated",
			Extensions: map[string]interface{}{
				"code":        "UNAUTHORIZED",
				"errorDetail": gc.Errors,
			},
		}
		return gErr
	} else {
		return nil
	}
}
