package bundler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"../config"
)

const DLL_WRITE_FUNCTION = `-- This function generated by LuBu. Used to load compressed .dll modules
local function LUBU_LOAD_COMPRESSED_DYNAMIC_LIBRARY(moduleName, bytes)
	local tempFileName = moduleName .. ".dll";
	local file = io.open(tempFileName, "wb");
	assert(file, '[LuBu Error] Error loading compressed dynamic library (counld not open file): ' .. tempFileName);
	for _, byte in ipairs(bytes) do
		file:write(string.char(byte));
	end
	file:close();
	local module = require(moduleName);
	os.remove(tempFileName);
	return module;
end`

func IsDllModuleDefined(cfg config.Config) bool {
	for _, file := range cfg.Modules {
		if strings.HasSuffix(file, ".dll") {
			return true
		}
	}
	return false
}

func GenerateDll(basePath, name, file string) string {
	fullPath := filepath.Join(basePath, file)
	log.Printf("Bundling \"%s\" module from: %s", name, fullPath)
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Error reading dll %s: %s", fullPath, err.Error())
	}
	byteAsLuaString := []string{}
	for _, _byte := range bytes {
		byteAsLuaString = append(byteAsLuaString, fmt.Sprintf("%d", _byte))
	}

	return fmt.Sprintf("\n-- Module \"%s\" (from %s)\npackage.preload['%s'] = LUBU_LOAD_COMPRESSED_DYNAMIC_LIBRARY('%s', { %s });", name, fullPath, name, name, strings.Join(byteAsLuaString, ", "))
}
