package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// use screen name rather than email for the security risk reason
type jwtCustomClaims struct {
	ScreenName string `json:"screen_name"`
	jwt.RegisteredClaims
}

type jwtService struct {
	issuer string
}

func (j *jwtService) GenerateToken(screenName string) (string, error) {
	tokenLifeSpan, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXP_HOUR"))
	if err != nil {
		return "", err
	}

	claims := &jwtCustomClaims{
		screenName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(tokenLifeSpan))),
			Issuer:    j.issuer,
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

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
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

// receive soon-to-be-expired token and return new token, for let the user stay logged in
func (j *jwtService) RefreshToken(tokenString string) (string, error) {
	token, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	screenName := claims["screen_name"].(string)

	return j.GenerateToken(screenName)
}
