package config

type PostgresDBConfig struct {
	Host               string `mapstructure:"db_host"`
	Port               int    `mapstructure:"db_port"`
	User               string `mapstructure:"db_user"`
	Password           string `mapstructure:"db_password"`
	DBName             string `mapstructure:"db_dbname"`
	SSLMode            string `mapstructure:"db_ssl_mode"`
	MaxOpenConnections int    `mapstructure:"db_max_open_connections"`
	MaxIdleConnections int    `mapstructure:"db_max_idle_connections"`
}
