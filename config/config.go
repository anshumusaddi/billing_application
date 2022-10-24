package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() error {
	viper.SetConfigType("yaml")
	configFile := os.Getenv("ENV")
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
