package bundler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GenerateLua(basePath, name, file string, isMainFile bool) string {
	fullPath := filepath.Join(basePath, file)
	log.Printf("Bundling \"%s\" module from: %s", name, fullPath)
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Error reading module \"%s\" (%s):\n\t%v", name, file, err)
	}
	if isMainFile {
		return fmt.Sprintf(INIT_PATTERN, fullPath, string(bytes))
	}
	return fmt.Sprintf(MODULE_PATTERN, name, fullPath, name, string(bytes))
}
