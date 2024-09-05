package auth

type Config struct {
	Token string `toml:"token"`
}

func NewConfig() *Config {
	return &Config {
	}
}
