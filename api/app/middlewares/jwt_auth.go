package middlewares

import (
	"app/ent"
	"app/ent/user"
	"app/services"
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
// https://gqlgen.com/recipes/authentication/
var authedCheckCtxKey = &contextKey{"isAuthed"}

type contextKey struct {
	isAuthed string
}

func JWTAuth(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// add gin.Context to context.Context so the resolvers can access the gin.Context in order to retrieve the isAuthed value
		// https://gqlgen.com/recipes/gin/
		ctx.Set(authedCheckCtxKey.isAuthed, false)
		c := context.WithValue(ctx.Request.Context(), "GinContextKey", ctx)
		ctx.Request = ctx.Request.WithContext(c)

		// Bearer token will be shown like `Authorization: Bearer <token>` in http header
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			log.Println(errors.New("auth header is invalid"))
			ctx.Next()
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := services.New(client).ValidateToken(tokenString)

		if err != nil {
			log.Println("authenticating request: ", err)
			ctx.Next()
		}

		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)

		// check if user exists in db
		_, err = client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
		if err != nil {
			log.Println("authenicating request: ", err)
			ctx.Next()
		}

		// token verified, user exists; thus this request is authorized; overwrite context value
		ctx.Set(authedCheckCtxKey.isAuthed, true)
		c = context.WithValue(ctx.Request.Context(), "GinContextKey", ctx)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}

// ForContext finds the isAuthed value from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) bool {
	raw := ctx.Value(authedCheckCtxKey.isAuthed)
	return raw.(bool)
}
