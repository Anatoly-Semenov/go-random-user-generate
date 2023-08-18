package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
	"time"
)

type Env string

const (
	PROD  Env = "prod"
	DEV   Env = "dev"
	LOCAL Env = "local"
)

type Config struct {
	Env        Env `yaml:"env" envDefault:"local"`
	HttpsSerer `yaml:"https_server"`
	Database   `yaml:"database"`
}

type HttpsSerer struct {
	Address       string        `yaml:"address" env-default:"localhost:8080"`
	Timeout       time.Duration `yaml:"timeout" env-default:"4s"`
	Iddle_timeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

type Database struct {
	Dialect  string `yaml:"dialect" env-default:"postgres"`
	Host     string `yaml:"host" env-default:"127.0.0.1"`
	Password string `yaml:"password" env-default:"random_users_pass"`
	User     string `yaml:"user" env-default:"random_users"`
	Db       string `yaml:"db" env-default:"random_users_db"`
	Port     string `yaml:"port" env-default:"5432"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	configPath := "../config/local.yaml"

	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			log.Println(help)
			log.Fatalf("Cannot read config: %s", err)
		}
	})

	return instance
}
