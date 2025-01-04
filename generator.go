package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const dllLoaderFunction = `
-- This function generated by LuBu. Used to load compressed .dll modules
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

func Generate(basePath string, config Config) []string {
	lines := []string{
		fmt.Sprintf("--[[\n%s\n]]", logo),
		"\n-- Constants",
		"LUBU_BUNDLED = true;",
		fmt.Sprintf("LUBU_BUILD_DATE = %d;", time.Now().Unix()),
	}

	// Constants
	for constName, constValue := range config.Const {
		switch constValue.(type) {
		case string:
			constValue = fmt.Sprintf("%q", constValue)
		case bool, float64:
			constValue = fmt.Sprintf("%v", constValue)
		default:
			log.Fatalf("Error, unsupported constant type: %T (const \"%s\")", constValue, constName)
		}
		lines = append(lines, fmt.Sprintf("%s = %s;", constName, constValue))
	}

	// Modules
	dllFunction := false
	for moduleName, fileName := range config.Modules {
		filePath := basePath + fileName
		log.Printf("Generating module \"%s\" from \"%s\"", moduleName, filePath)
		if strings.HasSuffix(fileName, ".lua") {
			lines = append(lines, fmt.Sprintf("\n--[[ Module \"%s\" (%s) ]]\npackage.preload['%s'] = (function()\n%s\nend)", moduleName, filePath, moduleName, ConvertLua(filePath)))
		} else if strings.HasSuffix(fileName, ".dll") {
			if !dllFunction {
				dllFunction = true
				lines = append(lines, dllLoaderFunction)
			}
			lines = append(lines, fmt.Sprintf("\n--[[ Module \"%s\" (%s) ]]\npackage.preload['%s'] = LUBU_LOAD_COMPRESSED_DYNAMIC_LIBRARY('%s', { %s })", moduleName, filePath, moduleName, moduleName, ConvertDll(filePath)))
		}
	}

	// Main file
	lines = append(lines, fmt.Sprintf("--[[ Entry Point (main file: \"%s\") ]]\nLUBU_BUNDLED_ENTRY_POINT = (function()\n%s\nend);\nLUBU_BUNDLED_ENTRY_POINT();", basePath+config.Main, ConvertLua(basePath+config.Main)))
	return lines
}

func ConvertLua(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading module file: %s\n", err.Error())
	}
	return string(bytes)
}

func ConvertDll(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading dll %s: %s", filePath, err.Error())
	}
	byteAsLuaString := []string{}
	for _, _byte := range bytes {
		byteAsLuaString = append(byteAsLuaString, fmt.Sprintf("%d", _byte))
	}
	return strings.Join(byteAsLuaString, ", ")
}
