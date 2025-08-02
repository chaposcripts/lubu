package bundler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/chaposcripts/lubu/config"
)

var luaKeywords = []string{"and",
	"break",
	"do",
	"else",
	"elseif",
	"end",
	"false",
	"for",
	"function",
	"if",
	"in",
	"local",
	"nil",
	"not",
	"or",
	"repeat",
	"return",
	"then",
	"true",
	"until",
	"while",
	"goto",
}

func checkVariableName(name string) bool {
	for _, keyword := range luaKeywords {
		if name == keyword {
			return false
		}
	}
	for index, char := range name {
		if (index == 0 && unicode.IsDigit(char)) && !unicode.IsLetter(char) && char != '_' {
			return false
		}
	}
	return true
}

func GenerateConstants(cfg config.Config) string {
	log.Println("Generating constants...")
	constantsCode := []string{
		"LUBU_BUNDLED = true;",
		fmt.Sprintf("LUBU_BUNDLED_AT = %s;", strconv.Itoa(int(time.Now().Unix()))),
	}
	for name, value := range cfg.Constants {
		if !checkVariableName(name) {
			log.Fatalf("Error, invalid variable name - \"%s\"", name)
		}

		switch value.(type) {
		case string:
			value = fmt.Sprintf("\"%s\"", value)
		case bool, float64:
			value = fmt.Sprintf("\"%v\"", value)
		default:
			log.Fatalf("Error, unsupported constant type: %T (const %s)", value, name)
		}
		constantsCode = append(constantsCode, fmt.Sprintf("%s = %s;", name, value))
		log.Printf("Created constant \"%s\" = %s", name, value)
	}

	constantsCodeString := "\n-- Constants\n" + strings.Join(constantsCode, "\n")
	if cfg.PrepareForObfuscation {
		constantsCodeString = PrepareForObfuscation(constantsCodeString)
	}
	return constantsCodeString
}
