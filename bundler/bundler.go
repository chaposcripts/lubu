package bundler

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"../config"
)

func Generate(basePath string, cfg config.Config) string {
	dllFunc := ""
	if IsDllModuleDefined(cfg) {
		dllFunc = DLL_WRITE_FUNCTION
	}
	items := []string{
		`--[[
	Bundled Using LuBu - Simple Lua Bundler
	LuBu: https://github.com/chaposcripts/lubu
]]`,
		dllFunc,
		GenerateConstants(cfg),
		GenerateModules(basePath, cfg),
	}

	return strings.Join(items, "\n")
}

func Bundle(basePath string, cfg config.Config) {
	outDir, _ := filepath.Split(cfg.Out)
	log.Printf("Writing code to \"%s\"\ndir: %s", cfg.Out, outDir)

	err := os.MkdirAll(outDir, 0644)
	if err != nil {
		log.Fatalf("Error creating directories for output file: %s", err.Error())
	}
	err = os.WriteFile(cfg.Out, []byte(Generate(basePath, cfg)), 0644)
	if err != nil {
		log.Fatalf("Error writing data: %v", err)
	}
	log.Printf("Done! Output file: %s", cfg.Out)

}
