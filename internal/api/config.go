package api

type Config struct {
	Addr string `toml:"addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		Addr: ":8080",
		LogLevel: "debug",
	}
}
