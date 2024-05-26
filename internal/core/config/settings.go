package config

import (
	"github.com/caarlos0/env/v7"
	"log"
	"sync"
)

type Config struct {
	AppInfo struct {
		AppName string `env:"APP_NAME" envDefault:"RAT"`
	}
	Logger struct {
		ErrorLvl bool `env:"RAT_ERROR_LOGGER_LVL" envDefault:"false"`
	}
}

var config *Config
var once sync.Once

func Init() {
	once.Do(func() {
		log.SetFlags(0)
		config = &Config{}
		err := env.Parse(config)
		if err != nil {
			log.Fatalf("failed to parse config: %v", err)
		}
	})
}

func Get() *Config {
	return config
}
