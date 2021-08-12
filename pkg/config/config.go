package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Load(prefix string, cfg interface{}) error {
	err := envconfig.Process(prefix, cfg)
	if err != nil {
		return err
	}
	return nil
}

func LoadFromFile(prefix string, cfg interface{}, filePath string) error {
	err := godotenv.Load(filePath)
	if err != nil {
		return err
	}
	return Load(prefix, cfg)
}
