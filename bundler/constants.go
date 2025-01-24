package bundler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"../config"
)

func GenerateConstants(cfg config.Config) string {
	log.Println("Generating constants...")
	constantsCode := []string{
		"LUBU_BUNDLED = true;",
		fmt.Sprintf("LUBU_BUNDLED_AT = %s;", strconv.Itoa(int(time.Now().Unix()))),
	}
	for name, value := range cfg.Constants {
		switch value.(type) {
		case string:
			constantsCode = append(constantsCode, fmt.Sprintf("%s = \"%s\";", name, value))
		case bool, float64:
			constantsCode = append(constantsCode, fmt.Sprintf("%s = %v;", name, value))
		default:
			log.Fatalf("Error, unsupported constant type: %T (const %s)", value, name)
		}
	}

	return "\n-- Constants\n" + strings.Join(constantsCode, "\n")
}
