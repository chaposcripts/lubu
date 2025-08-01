package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/chaposcripts/lubu/bundler"
	"github.com/chaposcripts/lubu/config"
	"github.com/chaposcripts/lubu/watcher"
)

func main() {
	log.Println("LuBu Started!")
	log.Println("Thanks for using LuBu! GitHub: https://github.com/chaposcripts/lubu/")
	if len(os.Args) < 2 {
		log.Fatalf("Error, config file not found!")
	}
	absoluteConfigPath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("pizda")
	}
	basePath, _ := filepath.Split(absoluteConfigPath)
	cfgPath := absoluteConfigPath

	if !filepath.IsAbs(cfgPath) {
		cfgPath, _ = filepath.Abs(cfgPath)
	}

	// Read config
	cfg, err := config.Read(cfgPath)
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	if cfg.Main == "" {
		log.Fatalf("Error, main file is empty!")
	}
	if cfg.Out == "" {
		log.Fatalf("Error, output file is empty!")
	}

	// Validate output file path
	if !filepath.IsAbs(cfg.Out) {
		cfg.Out = filepath.Join(basePath, cfg.Out)
	}

	// Print used paths
	log.Println("Base Path:", basePath)
	log.Println("Config Path:", cfgPath)
	log.Println("Output Path:", cfg.Out)
	if cfg.PrepareForObfuscation {
		log.Println("Prepare for obfuscation is ENABLED!")
	}

	bundler.Bundle(basePath, cfg)
	if cfg.WatcherDelay > 0 {
		log.Println("Watcher started! Waiting for files change...")
		watcher.StartWatcher(basePath, cfg, absoluteConfigPath)
	} else {
		log.Println("Watcher disabled. \"watcher_delay\" must be more than zero!")
	}
}
