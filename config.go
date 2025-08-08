package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
)

type Config struct {
	ThemeName string `toml:"theme"`
}

type ConfigStore struct {
	configPath string
}

func DefaultConfig() Config {
	return Config{
		ThemeName: "default",
	}
}

func NewConfigStore() (*ConfigStore, error) {
	configFilePath, err := xdg.ConfigFile("cadenza/config.toml")
	if err != nil {
		return nil, fmt.Errorf("Failed to resolve config path: %w", err)
	}

	return &ConfigStore{
		configPath: configFilePath,
	}, nil
}

func (s *ConfigStore) Config() (Config, error) {
	_, err := os.Stat(s.configPath)
	if os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	dir, fileName := filepath.Split(s.configPath)
	if len(dir) == 0 {
		dir = "."
	}

	buf, err := fs.ReadFile(os.DirFS(dir), fileName)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to read config file: %w", err)
	}

	if len(buf) == 0 {
		return DefaultConfig(), nil
	}

	cfg := Config{}
	if err := toml.Unmarshal(buf, &cfg); err != nil {
		return Config{}, fmt.Errorf("Config file does not have a valid format: %w", err)
	}

	return cfg, nil
}
