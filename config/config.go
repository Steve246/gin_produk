package config

import "os"

type ApiConfig struct {
	Url string
}

type Config struct {
	ApiConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	//localhost:8888
	c.ApiConfig = ApiConfig{Url: api}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
