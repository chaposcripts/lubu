package bundler

import (
	"fmt"
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
			modulesCode = append(modulesCode, GenerateModule(basePath, name, file, false))
		} else {
			modulesCode = append(modulesCode, GenerateDll(basePath, name, file))
		}
	}
	modulesCode = append(modulesCode, GenerateModule(basePath, "init", cfg.Main, true))
	return strings.Join(modulesCode, "\n")
}

func GenerateModule(basePath, name, file string, isInit bool) string {
	fullPath := filepath.Join(basePath, file)
	log.Printf("Bundling \"%s\" module from: %s", name, fullPath)
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Error reading module \"%s\" (%s):\n\t%v", name, file, err)
	}
	if isInit {
		return fmt.Sprintf(INIT_PATTERN, fullPath, string(bytes))
	}
	return fmt.Sprintf(MODULE_PATTERN, name, fullPath, name, string(bytes))
}
