package mock

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func MockJWT() (string, error) {
	expirateTime := time.Now().Add(5 * time.Minute).Unix()

	tokenClaims := Claims{
		Username: "AdminOne",
		Email:    "admin@gmail.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirateTime,
			Issuer:    "admin@gmail.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}
