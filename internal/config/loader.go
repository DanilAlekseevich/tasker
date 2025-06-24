package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	v.AutomaticEnv()
	v.SetEnvPrefix("TASKER")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetDefault("db.sslmode", "disable")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка чтения конфига: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка разбора конфига: %w", err)
	}

	return &cfg, nil
}
