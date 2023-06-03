package middlewares

import (
	"app/ent"
	"app/ent/user"
	"app/services"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
// https://gqlgen.com/recipes/authentication/
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func AuthorizeJWT(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bearer token will be shown like `Authorization: Bearer <token>` in http header
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			log.Println(errors.New("auth header is invalid"))
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := services.NewJWTService().ValidateToken(tokenString)

		if err != nil {
			log.Println(err)
			ctx.Redirect(http.StatusFound, "/login")
			ctx.AbortWithError(http.StatusForbidden, err)
		}

		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)

		// check if user exists in db
		user, err := client.User.Query().Where(user.EmailEQ(email)).Only(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/login")
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}

		ctx.Set(userCtxKey.name, user)
	}
}
