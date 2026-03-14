// Package config loads loglint configuration from the working directory.
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const configFileName = ".loglint.yml"

// Config defines the supported loglint configuration fields.
type Config struct {
	DisabledRules []string `yaml:"disabledRules"`
}

// New loads configuration from .loglint.yml in the current working directory.
func New() (*Config, error) {
	path, err := findConfigFile()
	if err != nil {
		return &Config{}, err
	}

	if path == "" {
		return &Config{}, nil
	}

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return &Config{}, fmt.Errorf("read %s: %w", configFileName, err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return &Config{}, fmt.Errorf("parse %s: %w", configFileName, err)
	}

	return &cfg, nil
}

func findConfigFile() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}

	path := filepath.Join(wd, configFileName)

	_, err = os.Stat(path)
	if err == nil {
		return path, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	}

	return "", fmt.Errorf("stat %s: %w", path, err)
}
