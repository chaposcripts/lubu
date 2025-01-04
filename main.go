package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const logo = `
	_           ____        
	| |         |  _ \       
	| |    _   _| |_) |_   _ 
	| |   | | | |  _ <| | | |
	| |___| |_| | |_) | |_| |
	|______\__,_|____/ \__,_|
	
	https://github.com/chaposcripts/lubu	
`

type Config struct {
	Modules map[string]string      `json:"modules"`
	Main    string                 `json:"main"`
	Output  string                 `json:"output"`
	Const   map[string]interface{} `json:"const"`
}

func main() {
	fmt.Print(logo)
	log.SetPrefix("LuBu ")
	if len(os.Args) < 2 {
		log.Panicln("Error, no input file")
	}
	jsonAbsolutePath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	basePath := filepath.Dir(jsonAbsolutePath)
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err.Error())
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("Error reading JSON: %s\n", err.Error())
	}
	if config.Main == "" || config.Output == "" {
		log.Fatalf("Error, field \"main\" or \"output\" is empty\n")
	}

	log.Printf("Base path: \"%s\"", basePath)

	lines := Generate(basePath, config)
	outFilePath := basePath + config.Output
	err = os.WriteFile(outFilePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing output file: %s\n", err.Error())
	}
	log.Printf("Done! Saved to \"%s\"\n", outFilePath)
}
