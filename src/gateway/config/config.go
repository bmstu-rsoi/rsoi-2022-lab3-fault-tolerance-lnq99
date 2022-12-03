package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	BonusUrl  string `yaml:"bonus_url"`
	FlightUrl string `yaml:"flight_url"`
	TicketUrl string `yaml:"ticket_url"`
}

type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type Config struct {
	Service ServiceConfig
	Server  ServerConfig
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

	return
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
