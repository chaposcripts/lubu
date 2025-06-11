package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Main         string                 `json:"main"`
	Out          string                 `json:"out"`
	Modules      map[string]string      `json:"modules"`
	WatcherDelay float64                `json:"watcher_delay"`
	Constants    map[string]interface{} `json:"const"`
}

const (
	OutputConfig = "autogen-config.json"
	SrcDir       = "./src"
)

var cfg = Config{
	Main:         "src/init.lua",
	WatcherDelay: 250,
	Out:          "dist/bundled.lua",
	Modules:      make(map[string]string),
	Constants: map[string]interface{}{
		"VERSION": "0.1-alpha",
	},
}

func collectLuaModules(rootDir string) (map[string]string, error) {
	modules := make(map[string]string)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %w", path, err)
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".lua" && ext != ".luac" && ext != ".dll" {
			return nil
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return fmt.Errorf("error getting relative path for %q: %w", path, err)
		}

		relPath = filepath.ToSlash(relPath)
		relPath = "src/" + relPath
		modulePath := strings.TrimSuffix(relPath, ext)
		moduleName := strings.ReplaceAll(modulePath, "/", ".")
		moduleName = strings.TrimPrefix(moduleName, "src.")

		modules[moduleName] = relPath
		return nil
	})

	return modules, err
}

func saveConfig() error {
	bytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	if err := os.WriteFile(OutputConfig, bytes, 0644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("Generating config...")

	modules, err := collectLuaModules(SrcDir)
	if err != nil {
		panic(err)
	}

	cfg.Modules = modules

	if err := saveConfig(); err != nil {
		panic(err)
	}

	fmt.Printf("Config successfully generated and saved to %s\n", OutputConfig)
}
