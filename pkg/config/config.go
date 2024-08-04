package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DatabaseConfig struct {
		User         string `json:"user"`
		Password     string `json:"password"`
		Address      string `json:"address"`
		DatabaseName string `json:"database_name"`
	} `json:"database_config"`
	ServerConfig struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"server_config"`
}

var c Config

func InitConfig(f string) {
	data, err := os.ReadFile(f)
	if err != nil {
		panic("unable to read config")
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		panic("unable to unmarshal config")
	}
}

func GetConfig() *Config {
	return &c
}
