package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env 		string 		`yaml:"env" env-required:"true"`
	StoragePath string 		`yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServer  `yaml:"http_server" env-required:"true"`
}

func LoadConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig("config/local.yaml", &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	return &cfg
}