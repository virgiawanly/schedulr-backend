package config

type ServerConfig struct {
	Ip   string `mapstructure:"server_ip"`
	Port int    `mapstructure:"server_port"`
}
