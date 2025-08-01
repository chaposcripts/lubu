package bundler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GenerateLua(basePath, name, file string, isMainFile, obfuscate bool) string {
	fullPath := filepath.Join(basePath, file)
	log.Printf("Bundling \"%s\" module from: %s", name, fullPath)
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Error reading module \"%s\" (%s):\n\t%v", name, file, err)
	}

	code := string(bytes)
	if obfuscate {
		code = PrepareForObfuscation(code)
		if isMainFile {
			code = strings.Replace(code, "function main()", "function __G.main()", 1)
		}
	}
	if isMainFile {
		return fmt.Sprintf(INIT_PATTERN, fullPath, code)
	}
	return fmt.Sprintf(MODULE_PATTERN, name, fullPath, name, code)
}
