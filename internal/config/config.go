package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type config struct {
	FilePath   string `yaml:"file_path"`
	DriverName string `yaml:"driver_name"`
}

func InitConfig() (*config, error) {
	yamlPath, _ := os.Getwd()
	yamlPath = filepath.Join(yamlPath, "..", "config.yaml")

	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}

	var cfg config

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
