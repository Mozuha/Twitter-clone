package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
include session id in the token so that we can invalidate the token by deleting the session
https://zenn.dev/ritou/articles/4a5d6597a5f250#%E3%80%8C%E3%82%BB%E3%83%83%E3%82%B7%E3%83%A7%E3%83%B3id%E3%82%92jwt%E3%81%AB%E5%86%85%E5%8C%85%E3%80%8D%E3%81%A8%E3%81%84%E3%81%86%E8%80%83%E3%81%88%E6%96%B9
*/
type jwtCustomClaims struct {
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

var issuer = "example_issuer"

func GenerateToken(userId string, sessionId string, forAccess bool) (string, error) {
	var (
		tokenLifeSpan int
		err           error
	)

	if forAccess {
		tokenLifeSpan, err = strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXP_HOUR"))
	} else {
		tokenLifeSpan, err = strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXP_HOUR"))
	}
	if err != nil {
		return "", err
	}

	claims := &jwtCustomClaims{
		userId,
		sessionId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(tokenLifeSpan))),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, err
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token.Valid {
		return token, nil
	} else {
		return nil, err
	}
}

// receive refresh token and return new access token, for let the user stay logged in
func RefreshToken(sessionId string, refTokenString string) (string, error) {
	token, err := ValidateToken(refTokenString)
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	uId := claims["user_id"].(string)

	// generate new access token
	return GenerateToken(uId, sessionId, true)
}
