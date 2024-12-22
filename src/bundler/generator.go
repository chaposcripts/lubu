package bundler

import (
	"strings"

	"../config"
)

func Generate(config config.Config, dir string) string {
	return strings.Join([]string{
		GenerateConstants(config),
		GenerateResources(config),
		GenerateDll(config),
		GenerateModules(dir, config),
	}, "\n\n")
}
