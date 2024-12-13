package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Main    string                 `json:"main"`
	Modules map[string]string      `json:"modules"`
	Output  string                 `json:"output"`
	Const   map[string]interface{} `json:"const"`
}

func ReadConfig(filePath string) Config {
	bytes, err := os.ReadFile(filePath + "/lubu.json")
	if err != nil {
		panic(fmt.Sprintf("Error reading lubu config file %s", err.Error()))
	}
	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(fmt.Sprintf("Error reading lubu config file %s", err.Error()))
	}
	return config
}
