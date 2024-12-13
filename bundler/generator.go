package bundler

import (
	"fmt"
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
		fmt.Printf("[LuBu][CONSTANT] Created constant \"%s\" = ", name)
		fmt.Println(value)
		constStringVal, isString := value.(string)
		if isString {
			value = "\"" + constStringVal + "\""
		}
		consts = append(consts, fmt.Sprintf("%s = %v;", name, value))
	}
	return strings.Join(consts, "\n")
}
