package config

import (
	"os"

	"github.com/skinkvi/crpt/pkg/util/logger"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"db"`
	Rabbitmq struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"rabbitmq"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, err
	}
	return cfg, nil
}

func GetConfig() *Config {
	return &Config{}
}
