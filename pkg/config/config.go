// config/config.go
package config

import (
	"github.com/spf13/viper"
)

// Config holds the configuration settings for the application.
type Config struct {
	Port                          int
	Environment                   string
	FirebaseServiceAccountKeyPath string
}

var AppConfig *Config

// LoadConfig loads the configuration settings from a config file.
func LoadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	AppConfig = &Config{
		Port:                          viper.GetInt("app.port"),
		Environment:                   viper.GetString("app.environment"),
		FirebaseServiceAccountKeyPath: viper.GetString("firebase.service_account_key_path"),
	}

	return nil
}
