package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port      int       `yaml:"port"`
	Services  []Service `yaml:"services"`
	RateLimit RateLimit `yaml:"rate_limit"`
}

type Service struct {
	Name         string   `yaml:"name"`
	Backends     []string `yaml:"backends"`
	AuthRequired bool     `yaml:"auth_required"` // <--- ADD THIS
}

type RateLimit struct {
	Requests   int `yaml:"requests"`
	PerSeconds int `yaml:"per_seconds"`
}

// LoadConfig reads YAML and converts it into Config struct
func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
