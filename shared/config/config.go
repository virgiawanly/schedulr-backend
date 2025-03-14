package config

import (
	"github.com/spf13/viper"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
)

func LoadConfig(cfg interface{}, defaults map[string]interface{}, envBindings map[string]string) error {
	v := viper.New()
	v.AutomaticEnv()

	for key, envVar := range envBindings {
		v.BindEnv(key, envVar)
	}

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	if err := v.Unmarshal(cfg); err != nil {
		logger.Errorf("Failed to unmarshal config: %v", err)
		return err
	}

	return nil
}
