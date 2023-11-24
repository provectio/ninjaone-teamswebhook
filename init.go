package main

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/provectio/godotenv"
)

func init() {
	if godotenv.Load() != nil {
		log.Info("No .env file, loading environnements variables")
	}

	if level, err := log.ParseLevel(os.Getenv("LOG_LEVEL")); err != nil {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(level)
	}

	secureToken = os.Getenv("SECURE_TOKEN")

	if env := os.Getenv("LISTEN_PORT"); env != "" {
		port = env
	}

	if env := os.Getenv("TEMPLATES_DIRECTORY"); env != "" {
		templatesDirectory = env
	}

	for _, fullEnv := range os.Environ() {
		if !strings.HasPrefix(fullEnv, "WEBHOOK_") {
			continue
		}
		split := strings.Split(fullEnv, "=")
		if len(split) < 2 {
			log.Fatal("Configuration environnement empty", "env", fullEnv)
		}
		name := strings.ToLower(strings.TrimPrefix(split[0], "WEBHOOK_"))
		url := strings.Join(split[1:], "=")
		webhooks[name] = url
		log.Debug("Webhook added", "name", name, "url", url)
	}

}
