package config

import (
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"db"`
	Rabbitmq struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"rabbitmq"`
}

func NewConfig(configPath string, logger *zap.Logger) (*Config, error) {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		logger.Error("Failed to read config file", zap.Error(err))
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		logger.Error("Failed to unmarshal config", zap.Error(err))
		return nil, err
	}

	return &cfg, nil
}
