package config

import (
	"os"
	"path/filepath"
	"strings"
)

func Generate(mainFile string, basePath string, outputFile string) error {
	var config Config
	err := filepath.Walk(basePath, func(file string, info os.FileInfo, err error) error {
		if strings.HasSuffix(file, ".lua") || strings.HasSuffix(file, ".dll") {
			_, moduleFile := filepath.Split(file)
			config.Modules[moduleFile] = file
		}
		return nil
	})
	return err
}

func ScanPath(path string) (files []string) {

	return
}
