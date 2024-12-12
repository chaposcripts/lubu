package main

import (
	"./bundler"
	"./config"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[1])
	config := config.ReadConfig(os.Args[1])
	err := os.MkdirAll(strings.Join(strings.Split(config.Output, "/")[:1], "/"), 0755)
	if err != nil {
		panic(fmt.Sprintf("Error creating output directory: %s", err.Error()))
	}

	err = os.WriteFile(config.Output, []byte(bundler.Generate(config)), 0644)
	if err != nil {
		panic(fmt.Sprintf("Error creating output file: %s", err.Error()))
	}
	fmt.Println("Done, saved to", config.Output)
}
