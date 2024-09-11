package config

// config struct for Viper Config
type Config struct {
	Port     string         `mapstructure:"port"`
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	Name    string `mapstructure:"name"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"password"`
	Address string `mapstructure:"address"`
}
