package telemetryApi

import (
	"fmt"
	"github.com/spf13/viper"
)

// APIConfiguration contains the different config fields.
type APIConfiguration struct {
	AppName     string `mapstructure:"APP_NAME"`
	AppPort     string `mapstructure:"APP_PORT"`
	PGDatabase  string `mapstructure:"PG_DB"`
	PGHost      string `mapstructure:"PG_HOST"`
	PGPort      int    `mapstructure:"PG_PORT"`
	PGUser      string `mapstructure:"PG_USER"`
	PGPassword  string `mapstructure:"PG_PASSWORD"`
	PGSchema    string `mapstructure:"PG_SCHEMA"`
	PGSSLMode   string `mapstructure:"PG_SSLMODE" default:"disable" validate:"oneof=verify-full verify-ca require prefer allow disable"`
	CorsEnabled bool   `mapstructure:"CORS_ENABLED"`
	CorsOrigins string `mapstructure:"CORS_ALLOWED_ORIGINS"`
}

// APIConfig is the variable containing all the configuration for the app to run.
var APIConfig APIConfiguration

// LoadConfig will read the config file and save the values in a Configuration struct.
func LoadConfig(path string) {
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s \n", err)
	}

	if err := viper.Unmarshal(&APIConfig); err != nil {
		fmt.Printf("Unable to decode into struct: %v \n", err)
		return
	}
}
