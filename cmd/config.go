package main

import (
	"strings"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/config"
	"github.com/spf13/viper"
)

func readConfig(configFile string) (*config.Config, error) {
	var config config.Config
	if configFile != "" {
		viper.SetConfigFile(configFile)
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
		err = viper.Unmarshal(&config)
		return &config, err
	}
	// Automatically look for environment variables
	viper.AutomaticEnv()

	// Replace dots with underscores to map env variables like DATABASE_NAME
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// You can explicitly bind environment variables to nested fields
	_ = viper.BindEnv("database.name", "DATABASE_NAME")
	_ = viper.BindEnv("database.user", "DATABASE_USER")
	_ = viper.BindEnv("database.password", "DATABASE_PASSWORD")
	_ = viper.BindEnv("database.address", "DATABASE_ADDRESS")
	_ = viper.BindEnv("port", "PORT")

	// Unmarshal the environment variables into the struct
	err := viper.Unmarshal(&config)
	return &config, err
}
