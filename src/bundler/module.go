package bundler

import (
	"fmt"
	"log"
	"os"
	"strings"

	"../config"
)

func GenerateLuaModule(name, file string, isMain bool) string {
	label := "MODULE"
	if isMain {
		label = "MAIN"
	}
	log.Printf("[%s] Bundling \"%s\" from \"%s\"\n", label, name, file)
	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading module %s (%s): %s", name, file, err.Error())
	}
	if isMain {
		return fmt.Sprintf("-- Entry Point %s (%s)\nLUBU_BUNDLED_ENTRY_POINT = (function(...)\n%s\nend)\nLUBU_BUNDLED_ENTRY_POINT();", name, file, string(bytes))
	}
	return fmt.Sprintf("-- Module %s (%s)\npackage.preload['%s'] = (function(...)\n%s\nend)", name, file, name, string(bytes))
}

func GenerateModules(dir string, config config.Config) string {
	result := []string{}
	for moduleName, modulePath := range config.Modules {
		if len(result) == 0 {
			result = append(result, "-- LuBu Bundled Lua Modules")
		}
		result = append(result, GenerateLuaModule(moduleName, dir+"\\"+modulePath, false)+"\n")
	}
	result = append(result, GenerateLuaModule("main", dir+"\\"+config.Main, true)+"\n")
	return strings.Join(result, "\n")
}
