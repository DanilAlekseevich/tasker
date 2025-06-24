package config

type Config struct {
	DB DatabaseConfig `mapstructure:"db"`
}
