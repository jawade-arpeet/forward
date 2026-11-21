package config

import (
	"fmt"
	"strings"

	"forward/internal/constants"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var cfg *Config

type ServerCfg struct {
	RunEnv constants.Env `mapstructure:"run_env" validate:"oneof=DEV STG PROD"`
	Port   int           `mapstructure:"port" validate:"required"`
}

type PostgresCfg struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Database string `mapstructure:"database" validate:"required"`
}

type RedisCfg struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password" validate:"required"`
	Database string `mapstructure:"database" validate:"required,numeric"`
}

type Config struct {
	Server   *ServerCfg   `mapstructure:"server" validate:"required"`
	Postgres *PostgresCfg `mapstructure:"postgres" validate:"required"`
	Redis    *RedisCfg    `mapstructure:"redis" validate:"required"`
}

func Load() error {
	v := viper.New()

	v.SetEnvPrefix("FORWARD")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	keys := []string{
		"server.run_env",
		"server.port",
		"postgres.host",
		"postgres.port",
		"postgres.username",
		"postgres.password",
		"postgres.database",
		"redis.host",
		"redis.port",
		"redis.username",
		"redis.password",
		"redis.database",
	}

	for _, key := range keys {
		if err := v.BindEnv(key); err != nil {
			return fmt.Errorf("failed to bind env %s: %w", key, err)
		}
	}

	loadedCfg := &Config{}
	if err := v.Unmarshal(loadedCfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validator.New().Struct(loadedCfg); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	cfg = loadedCfg
	return nil
}

func GetServerCfg() *ServerCfg {
	return cfg.Server
}

func GetPostgresCfg() *PostgresCfg {
	return cfg.Postgres
}

func GetRedisCfg() *RedisCfg {
	return cfg.Redis
}
