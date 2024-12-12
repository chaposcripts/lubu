package bundler

import (
	"fmt"
	"strings"

	"../config"
)

func Generate(config config.Config) string {
	result := []string{
		`--[[Bundled using LuBu - Simple Lua Bundler]]`,
	}
	for moduleName, modulePath := range config.Modules {
		fmt.Println("reading", moduleName, modulePath)
		result = append(result, ConvertFileAsPackage(moduleName, modulePath, false))
	}
	result = append(result, ConvertFileAsPackage("main", config.Main, true))
	return strings.Join(result, "\n\n")
}
