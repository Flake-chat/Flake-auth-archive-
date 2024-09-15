package api

import (
	"github.com/Flake-chat/Flake-auth/auth"
	"github.com/Flake-chat/Flake-auth/store"
)

type Config struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
	Auth     *auth.Config
}

func NewConfig() *Config {
	return &Config{
		LogLevel: "debug",
	}
}
