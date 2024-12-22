package bundler

import (
	"fmt"
	"log"
	"os"
	"strings"

	"../config"
)

/*

	Resource usage example (moonloader mimgui):

	local texture = RESOURCE_IMAGE_LOGO and imgui.CreateTextureFromFileInMemory(...) or imgui.CreatetextureFromFile(...)

	lubu.config:
	"resource": {
		"logo": { type:"PNG", "src/resource/logo.png" }
	}
*/

func ConvertToBase85(varName, filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[RESOURCE][ERROR] Error reading resource %s (%s): %s", varName, filePath, err.Error())
	}
	varName = fmt.Sprintf("-- Compressed Resource File %s (%s)\n%s = [[%s]];", varName, filePath, varName, strings.ReplaceAll(string(bytes), "\n", ""))
	return varName
}

func GenerateResources(config config.Config) string {
	list := []string{}
	for varName, path := range config.Resource {
		if len(list) == 0 {
			list = append(list, "-- Base85 Compressed Resources")
		}
		list = append(list, ConvertToBase85(varName, path))
	}
	return strings.Join(list, "\n")
}
