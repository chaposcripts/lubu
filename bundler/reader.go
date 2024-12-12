package bundler

import (
	"fmt"
	"os"
)

func ConvertFileAsPackage(name, file string, isMain bool) string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Error reading module %s (%s): %s", name, file, err.Error()))
	}
	if isMain {
		return fmt.Sprintf("--Entry Point %s (%s)\nlocal entry = (function(...)\n%s\nend)\nentry();", name, file, string(bytes))
	}
	return fmt.Sprintf("--Module %s (%s)\npackage.preload['%s'] = (function(...)\n%s\nend)", name, file, name, string(bytes))
}
