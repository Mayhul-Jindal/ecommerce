package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	DB_URL                 string        `mapstructure:"DB_URL"`
	SERVER_PORT            string        `mapstructure:"SERVER_PORT"`
	LOG_RETENTION_POLICY   string        `mapstructure:"LOG_RETENTION_POLICY"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	TOKEN_SYMMETRIC_KEY    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	RAZORPAY_KEY_ID        string        `mapstructure:"RAZORPAY_KEY_ID"`
	RAZORPAY_KEY_SECRET    string        `mapstructure:"RAZORPAY_KEY_SECRET"`
	EMAIL_SENDER_NAME      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EMAIL_SENDER_ADDRESS   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EMAIL_SENDER_PASSWORD  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
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
