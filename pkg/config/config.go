package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	Database         string `yaml:"database"`
	MaxAttempts      int    `yaml:"max_attempts"`
	SecondsToConnect int    `yaml:"seconds_to_connect"`
}

func LoadConfig(path string) (DatabaseConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return DatabaseConfig{}, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	var cfg DatabaseConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return DatabaseConfig{}, fmt.Errorf("ошибка парсинга yaml: %w", err)
	}

	return cfg, nil
}
