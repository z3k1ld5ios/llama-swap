package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mostlygeek/llama-swap/proxy"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var (
		configFile  = flag.String("config", "config.yaml", "path to configuration file")
		listenAddr  = flag.String("listen", "0.0.0.0:8080", "address to listen on")
		showVersion = flag.Bool("version", false, "show version information")
		logLevel    = flag.String("log-level", "info", "log level (debug, info, warn, error)")
	)

	flag.Parse()

	if *showVersion {
		fmt.Printf("llama-swap version %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	// Load configuration
	cfg, err := proxy.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Create and start the proxy server
	server, err := proxy.New(cfg, *listenAddr, *logLevel)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Printf("llama-swap %s starting on %s", version, *listenAddr)

	if err := server.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
