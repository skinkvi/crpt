package config

import (
	"os"

	"github.com/skinkvi/crpt/pkg/util/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Logger *zap.Logger
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

var cfg *Config

func NewConfig(configPath string) (*Config, error) {
	err := logger.InitLogger()
	if err != nil {
		return nil, err
	}

	cfg = &Config{
		Logger: logger.GetLogger(),
	}
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return nil, err
	}
	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}
