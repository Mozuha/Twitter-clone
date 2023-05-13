package middlewares

import (
	"api/services"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT() gin.HandlerFunc {
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

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Email]: ", claims["email"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
		}
	}
}
