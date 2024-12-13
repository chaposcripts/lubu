package bundler

import (
	"fmt"
	"log"
	"strings"

	"../config"
)

func Generate(config config.Config, dir string) string {
	result := []string{
		"LUBU_BUNDLED = true;",
		"--LuBu Constants\n" + GenerateConstants(config),
	}
	for moduleName, modulePath := range config.Modules {
		result = append(result, ConvertFileAsPackage(moduleName, dir+"\\"+modulePath, false))
	}
	result = append(result, ConvertFileAsPackage("main", dir+"\\"+config.Main, true))
	return strings.Join(result, "\n\n")
}

func GenerateConstants(config config.Config) string {
	consts := []string{}
	for name, value := range config.Const {
		_, isBool := value.(bool)
		_, isNumber := value.(float64)
		constAsString, isString := value.(string)
		if !isBool && !isNumber && !isString {
			log.Fatalf("[ERROR] [CONSTANT] Unsupported constant \"%s\" type: %T. Only string, number and bool are supported", name, value)
		}
		if isString {
			value = "\"" + constAsString + "\""
		}
		log.Printf("[CONSTANT] Created constant \"%s\" = %v", name, value)
		consts = append(consts, fmt.Sprintf("%s = %v;", name, value))
	}
	return strings.Join(consts, "\n")
}
