package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

type ConfigSchema struct {
	Server   ServerConfig   `json:"server"`
	Postgres PostgresConfig ` json:"postgres"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type PostgresConfig struct {
	Database        string `json:"database"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Port            int    `json:"port"`
	Host            string `json:"host"`
	MaxConn         int    `json:"max_conn"`
	MaxIdleConn     int    `json:"max_idle_conn"`
	MaxIdleConnTime int    `json:"max_idle_conn_time"`
}

func LoadConfig() (*ConfigSchema, error) {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(DefaultConfig)); err != nil {
		return nil, err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	cfg := ConfigSchema{}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
