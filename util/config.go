package util

import (

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	DB_URL string `mapstructure:"DB_URL"`
	SERVER_PORT string `mapstructure:"SERVER_PORT"`
	LOG_RETENTION_POLICY string `mapstructure:"LOG_RETENTION_POLICY"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
