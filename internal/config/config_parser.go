package config

import "os"

func CreateConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg := &Config{}
	cfg.Server.Port = port

	return cfg, nil
}
