package bundler

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"../config"
)

const MODULE_PATTERN = "\n-- Module \"%s\" (from %s)\npackage.preload['%s'] = (function()\n%s\nend);"
const INIT_PATTERN = "\n-- Init (from %s) \nLUBU_ENTRY_POINT = (function()\n%s\nend);\nLUBU_ENTRY_POINT();"

func GenerateModules(basePath string, cfg config.Config) string {
	log.Println("Generating modules...")
	modulesCode := []string{}
	for name, file := range cfg.Modules {
		if strings.HasSuffix(file, ".lua") {
			modulesCode = append(modulesCode, GenerateLua(basePath, name, file, false))
		} else {
			modulesCode = append(modulesCode, GenerateDll(basePath, name, file))
		}
	}
	modulesCode = append(modulesCode, GenerateLua(basePath, "init", cfg.Main, true))
	return strings.Join(modulesCode, "\n")
}

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
		GenerateResources(basePath, cfg),
		GenerateModules(basePath, cfg),
	}

	return strings.Join(items, "\n")
}

func Bundle(basePath string, cfg config.Config) {
	outDir, _ := filepath.Split(cfg.Out)
	log.Printf("Writing code to \"%s\"", cfg.Out)

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
