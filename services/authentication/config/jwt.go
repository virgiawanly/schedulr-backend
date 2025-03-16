package config

import (
	"time"

	sharedConfig "github.com/virgiawanly/schedulr-backend/shared/config"
)

type JWTConfig struct {
	SecretKey       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func (jwtConfig *JWTConfig) LoadConfig() error {
	defaults := map[string]interface{}{
		"jwt_secret_key":        "",
		"jwt_access_token_ttl":  "15m",
		"jwt_refresh_token_ttl": "60m",
	}

	envBindings := map[string]string{
		"jwt_secret_key":        "JWT_SECRET_KEY",
		"jwt_access_token_ttl":  "JWT_ACCESS_TOKEN_TTL",
		"jwt_refresh_token_ttl": "JWT_REFRESH_TOKEN_TTL",
	}

	return sharedConfig.LoadConfig(jwtConfig, defaults, envBindings)
}
