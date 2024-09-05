package store

type Config struct {
	DB string `toml:"db_url"`
}

func NewConfig() *Config {
	return &Config {
	}
}
