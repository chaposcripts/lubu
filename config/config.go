package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Main         string                 `json:"main"`
	Out          string                 `json:"out"`
	Modules      map[string]string      `json:"modules"`
	WatcherDelay float64                `json:"watcher_delay"`
	Constants    map[string]interface{} `json:"const"`
}

func Read(file string) (cfg Config, err error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &cfg)
	return
}
