package main

import (
	"log"
	"os"

	"fetch-test/internal/application"
	"fetch-test/internal/configuration"
)

func main() {
	config, err := configuration.GetConfig()
	if err != nil {
		log.Fatal("Invalid Configuration")
		os.Exit(1)
	}

	app, err := application.New(config)
	if err != nil {
		log.Fatal("Failed to initialize application")
		os.Exit(1)
	}
	app.Run()
}
