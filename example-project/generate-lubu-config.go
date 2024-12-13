package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Main    string                 `json:"main"`
	Modules map[string]string      `json:"modules"`
	Output  string                 `json:"output"`
	Const   map[string]interface{} `json:"const"`
}

var cfg = Config{
	Main:    "src/init.lua",
	Modules: map[string]string{},
	Output:  "dist/release.lua",
	Const:   map[string]interface{}{},
}
var ignoreDir = []string{"src\\resource\\cache"}

func main() {
	// lines := []string{
	// 	fmt.Sprintf("Main \"%s\"", cfg.mainFile),
	// }
	err := filepath.Walk("./src", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".lua") && !isDirIgnored(path) && path != cfg.Main {
			cfg.Modules[getFileModuleName(path)] = path
			// lines = append(lines, fmt.Sprintf("Module \"%s\" \"%s\"", getFileModuleName(path), strings.ReplaceAll(path, "\\", "/")))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// lines = append(lines, fmt.Sprintf("Output \"%s\"", outputFile))
	bytes, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	os.WriteFile("./lubu.json", []byte(bytes), 0644)
	fmt.Println("Done, saved to ./lubu.json")
}

func getFileModuleName(path string) string {
	parts := strings.Split(strings.ReplaceAll(path, ".lua", ""), "\\")
	return strings.Join(parts[1:], ".")
}

func isDirIgnored(dir string) bool {
	for _, dirName := range ignoreDir {
		if strings.HasPrefix(dir, dirName) {
			return true
		}
	}
	return false
}
