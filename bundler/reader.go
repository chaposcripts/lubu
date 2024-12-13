package bundler

import (
	"fmt"
	"log"
	"os"
)

func ConvertFileAsPackage(name, file string, isMain bool) string {
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
		return fmt.Sprintf("--Entry Point %s (%s)\nlocal entry = (function(...)\n%s\nend)\nentry();", name, file, string(bytes))
	}
	return fmt.Sprintf("--Module %s (%s)\npackage.preload['%s'] = (function(...)\n%s\nend)", name, file, name, string(bytes))
}
