package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Main     string                 `json:"main"`
	Modules  map[string]string      `json:"modules"`
	Output   string                 `json:"output"`
	Const    map[string]interface{} `json:"const"`
	Dll      map[string]string      `json:"dll"`
	Resource map[string]string      `json:"resource"`
}

func ReadConfig(filePath string) Config {
	bytes, err := os.ReadFile(filePath + "/lubu.json")
	if err != nil {
		log.Fatalf("Error reading lubu config file %s", err.Error())
	}
	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("[ERROR] Error reading lubu config file %s", err.Error())
	}
	return config
}
