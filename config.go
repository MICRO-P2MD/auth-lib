package auth

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AUTH_URL string `envconfig:"AUTH_URL" default:"http://localhost:7001"`
}

func NewConfig() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
