package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ApiUrl struct {
	Url string `yaml:"apiurl"`
}

func LoadAPIUrl(path string) (ApiUrl, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ApiUrl{}, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	var cfg ApiUrl
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return ApiUrl{}, fmt.Errorf("ошибка парсинга yaml: %w", err)
	}

	return cfg, nil
}
