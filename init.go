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

	var err error
	if logLevel, err = log.ParseLevel(os.Getenv("LOG_LEVEL")); err != nil {
		log.Warn("Invalid log level, fallback to info", "error", err)
	} else {
		log.SetLevel(logLevel)
	}

	secureToken = os.Getenv("SECURE_TOKEN")

	if env := os.Getenv("LISTEN_PORT"); env != "" {
		port = env
	}

	if env := os.Getenv("TEMPLATES_DIRECTORY"); env != "" {
		templatesDirectory = env
	}

	// Testing directory exist
	if _, err := os.Stat(templatesDirectory); os.IsNotExist(err) {
		log.Fatal("Templates directory not found", "directory", templatesDirectory)
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
