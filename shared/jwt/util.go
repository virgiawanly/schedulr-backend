package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewJWTConfig(secretKey string, accessTTL, refreshTTL time.Duration) *JWTConfig {
	return &JWTConfig{
		SecretKey:       secretKey,
		AccessTokenTTL:  accessTTL,
		RefreshTokenTTL: refreshTTL,
	}
}

func (cfg *JWTConfig) GenerateToken(accountID, role, email string, isRefreshToken bool) (string, error) {
	now := time.Now()

	var expiresAt *jwt.NumericDate
	if isRefreshToken {
		expiresAt = jwt.NewNumericDate(now.Add(cfg.RefreshTokenTTL))
	} else {
		expiresAt = jwt.NewNumericDate(now.Add(cfg.AccessTokenTTL))
	}

	claims := AccountClaims{
		AccountID: accountID,
		Role:      role,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.SecretKey))
}

func (cfg *JWTConfig) ValidateToken(tokenString string) (*AccountClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccountClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(cfg.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccountClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
