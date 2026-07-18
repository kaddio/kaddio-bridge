package main

import (
	"fmt"
	"log"
	"os"

	"kaddio-bridge/internal/api"
	"kaddio-bridge/internal/config"
)

const version = "0.1.5"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			fmt.Printf("kaddio-bridge %s\n", version)
			return
		case "token":
			cfg, err := config.Load()
			if err != nil {
				log.Fatalf("Error loading config: %v\n", err)
			}
			fmt.Println(cfg.Token)
			return
		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
			fmt.Fprintf(os.Stderr, "Usage: kaddio-bridge [version|token]\n")
			os.Exit(1)
		}
	}

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
