package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultFile string `json:"default_file"`
}

func configPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(configDir, "garden", "config.json")
}

func LoadOrCreateConfig() Config {
	path := configPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), 0755)
		defaultCfg := Config{}
		saveConfig(defaultCfg)
		return defaultCfg
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg
}

func saveConfig(cfg Config) {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(configPath(), data, 0644)
}
