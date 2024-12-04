package telemetryIngest

import (
	"github.com/spf13/viper"
	"log"
)

// ClientConfiguration contains the client configuration
type ClientConfiguration struct {
	ClientPort string `mapstructure:"CLIENT_PORT"`
	PGDatabase string `mapstructure:"PG_DB"`
	PGHost     string `mapstructure:"PG_HOST"`
	PGPort     int    `mapstructure:"PG_PORT"`
	PGUser     string `mapstructure:"PG_USER"`
	PGPassword string `mapstructure:"PG_PASSWORD"`
	PGSchema   string `mapstructure:"PG_SCHEMA"`
	PGSSLMode  string `mapstructure:"PG_SSLMODE" default:"disable" validate:"oneof=verify-full verify-ca require prefer allow disable"`
}

// ClientConfig is the variable containing all the configuration for the client application.
var ClientConfig ClientConfiguration

// LoadConfig is the function in charge of reading the configuration file or environment variables
func LoadConfig(path string) {
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("[ERROR] Error reading config file: %s \n", err)
	}

	if err := viper.Unmarshal(&ClientConfig); err != nil {
		log.Printf("[ERROR] Unable to decode into struct: %v \n", err)
		return
	}
}
