package watcher

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chaposcripts/lubu/bundler"
	"github.com/chaposcripts/lubu/config"
)

var lastCheckedTime map[string]time.Time = map[string]time.Time{}

func StartWatcher(basePath string, cfg config.Config, configFile string) {
	for range time.NewTicker(time.Millisecond * time.Duration(cfg.WatcherDelay)).C {
		hasChanged, file := checkAllFiles(basePath, cfg)
		if hasChanged || checkFile(cfg.Main, basePath) {
			log.Printf("Watcher: file \"%s\" was changed, re-bundling...", file)
			bundler.Bundle(basePath, cfg)
		}
	}
}

func checkAllFiles(basePath string, cfg config.Config) (bool, string) {
	for _, file := range cfg.Modules {
		return checkFile(file, basePath), file
	}
	return false, ""
}

func checkFile(file, basePath string) bool {
	filePath := file
	if !filepath.IsAbs(filePath) {
		filePath = filepath.Join(basePath, file)
	}
	fileStat, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Watcher: Error getting file modification date:\n\tFile: %s\n\tError: %v", filePath, err)
	}
	modTime := fileStat.ModTime()
	oldModTime, oldModTimeExists := lastCheckedTime[filePath]
	if !oldModTimeExists || (modTime != oldModTime && oldModTime.Unix() > 0) {
		lastCheckedTime[filePath] = modTime
	}
	return modTime != oldModTime && oldModTime.Unix() > 0
}
