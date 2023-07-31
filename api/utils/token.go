package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Token struct {
	User LoginClaim
	Iat  int64
	Exp  int64
}

type LoginClaim struct {
	AccountID   uint32
	AccountName string
	Email       string
}

func GetToken(c echo.Context) *jwt.Token {
	getToken := c.Get("Token")
	token := getToken.(*jwt.Token)
	return token
}

const (
	ExpireDate = 7
)

func GenerateToken(claim LoginClaim) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id":   claim.AccountID,
		"account_name": claim.AccountName,
		"email":        claim.Email,
		"iat":          now.Unix(),
		"exp":          now.Add(time.Hour * 24 * ExpireDate).Unix(),
	})

	return token.SignedString([]byte(secret))
}
