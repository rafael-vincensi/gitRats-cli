package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Token  string `json:"token"`
	APIURL string `json:"api_url"`
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error obtaining home directory: %w", err)
	}

	dir := filepath.Join(home, ".gitrats")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(dir, "config.json"), nil
}

func Save(cfg Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func Load() (Config, error) {
	var cfg Config

	path, err := configPath()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cfg.APIURL = "http://localhost:8080"
			return cfg, nil
		}
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	if cfg.APIURL == "" {
		cfg.APIURL = "http://localhost:8080"
	}

	return cfg, nil
}
