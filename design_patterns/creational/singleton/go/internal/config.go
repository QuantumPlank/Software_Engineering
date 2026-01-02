package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	AppName string
	Version string
}

var cfg *Config
var once sync.Once

func (cfg *Config) load() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(cfg)
}

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		err := cfg.load()
		if err != nil {
			log.Println("Error loading config:", err.Error())
			os.Exit(1)
		}
	})
	return cfg
}
