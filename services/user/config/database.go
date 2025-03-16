package config

import (
	"fmt"
	"strconv"

	sharedConfig "github.com/virgiawanly/schedulr-backend/shared/config"
)

type DBConfig struct {
	sharedConfig.PostgresDBConfig `mapstructure:",squash"`
}

func (db *DBConfig) LoadConfig() error {
	defaults := map[string]interface{}{
		"db_host":                 "localhost",
		"db_port":                 5432,
		"db_user":                 "postgres",
		"db_password":             "",
		"db_dbname":               "schedulr_user",
		"db_ssl_mode":             "disable",
		"db_max_open_connections": 25,
		"db_max_idle_connections": 25,
	}

	envBindings := map[string]string{
		"db_host":                 "DB_HOST",
		"db_port":                 "DB_PORT",
		"db_user":                 "DB_USER",
		"db_password":             "DB_PASSWORD",
		"db_dbname":               "DB_NAME",
		"db_ssl_mode":             "DB_SSL_MODE",
		"db_max_open_connections": "DB_MAX_OPEN_CONNECTIONS",
		"db_max_idle_connections": "DB_MAX_IDLE_CONNECTIONS",
	}

	return sharedConfig.LoadConfig(db, defaults, envBindings)
}

func (db *DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		db.User,
		db.Password,
		db.Host,
		strconv.Itoa(db.Port),
		db.DBName,
		db.SSLMode,
	)
}
