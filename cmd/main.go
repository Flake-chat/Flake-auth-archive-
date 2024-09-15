package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Flake-chat/Flake-auth/internal/api"
)

var (
	confPath string
	addr     string
)

func init() {
	flag.StringVar(&confPath, "config-path", "config/server.toml", "Path conf")
	flag.StringVar(&addr, "address", ":8080", "Address Port")

}

func main() {
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(confPath, config)

	if err != nil {
		log.Fatal(err)
	}
	config.Addr = addr
	s := api.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
