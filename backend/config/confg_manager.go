package config

import (
	"encoding/json"
	"os"
)

type DBConfig struct {
	Conn string `json:"conn"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type Config struct {
	Env    string       `json:"env"`
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

var config *Config

func GetConfig() *Config {
	return config
}

// Load the config from config.json
func LoadConfig() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}
	config = &cfg
	return nil
}
