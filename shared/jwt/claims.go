package jwt

import "github.com/golang-jwt/jwt/v5"

type AccountClaims struct {
	AccountID string `json:"account_id"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}
