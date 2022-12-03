package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Url string `mapstructure:"DB_URL"`
}

type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type Config struct {
	Db     DbConfig
	Server ServerConfig
}

func LoadConfig() (c *Config, err error) {
	filename := flag.String("configFile", "config.yaml", "Config file (default: config.yaml)")

	data, err := os.ReadFile(*filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return
	}

	c.Server.Host = getEnv("HOST", c.Server.Host)
	c.Server.Port = getEnv("PORT", c.Server.Port)
	c.Db.Url = getEnv("DB_URL", c.Db.Url)

	return
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
