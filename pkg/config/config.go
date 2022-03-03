package config

import (
	"bytes"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type ConfigSchema struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig ` json:"database"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DatabaseConfig struct {
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
	v := viper.New()
	v.SetEnvPrefix("wager")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()

	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBuffer(DefaultConfig)); err != nil {
		return nil, err
	}

	cfg := ConfigSchema{}
	if err := v.Unmarshal(&cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
	}); err != nil {
		return nil, err
	}
	return &cfg, nil
}
