package bundler

import (
	"fmt"
	"log"
	"strings"

	"../config"
)

func GenerateConstants(config config.Config) string {
	consts := []string{"-- LuBu Constants\nLUBU_BUNDLED = true;\n"}
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

func JsonToLuaTable(json string) string {

	return json
}
