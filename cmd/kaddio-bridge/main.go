package main

import (
	"log"

	"kaddio-bridge/internal/api"
	"kaddio-bridge/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	cfgPath, _ := config.Dir()
	log.Printf("Kaddio Bridge started\n\nConfig:\n%s/config.json\n\nToken:\n%s\n\n", cfgPath, cfg.Token)

	srv := api.New(cfg)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
