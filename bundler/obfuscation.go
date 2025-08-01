package bundler

import (
	"fmt"
	"regexp"
	"strings"
)

var GLOBAL_TABLES []string = []string{
	"math",
	"lua_thread",
}

func PrepareForObfuscation(code string) string {
	code = replaceFunctionDefinitions(code)
	code = replaceFunctionCalls(code)
	code = replaceNumbers(code)
	return code
}

func replaceNumbers(code string) string {
	processed := make(map[int]bool)

	protectedRegex := regexp.MustCompile(`(?s)(\[[[:space:]]*["'].*?["'][[:space:]]*\])|(".*?")|('.*?')`)
	protectedAreas := protectedRegex.FindAllStringIndex(code, -1)

	isProtected := func(pos int) bool {
		for _, area := range protectedAreas {
			if pos >= area[0] && pos < area[1] {
				return true
			}
		}
		return false
	}

	numberRegex := regexp.MustCompile(`\b[-+]?\d*\.?\d+\b`)
	tonumberRegex := regexp.MustCompile(`tonumber\("[^"]*"\)`)
	tonumberMatches := tonumberRegex.FindAllStringIndex(code, -1)

	var result strings.Builder
	lastPos := 0

	numberMatches := numberRegex.FindAllStringSubmatchIndex(code, -1)
	for _, match := range numberMatches {
		start, end := match[0], match[1]
		result.WriteString(code[lastPos:start])

		skip := false
		for _, tm := range tonumberMatches {
			if start >= tm[0] && end <= tm[1] {
				skip = true
				break
			}
		}

		if !skip && !isProtected(start) && !processed[start] {
			number := code[start:end]
			result.WriteString(fmt.Sprintf(`tonumber("%s")`, number))
			processed[start] = true
		} else {
			result.WriteString(code[start:end])
		}
		lastPos = end
	}
	result.WriteString(code[lastPos:])
	return result.String()
}

func replaceFunctionDefinitions(code string) string {
	code = regexp.MustCompile(`function\s+([\w_]+(?:\.[\w_]+)+)\s*\((.*?)\)([\s\S]*?)end`).
		ReplaceAllStringFunc(code, func(match string) string {
			parts := regexp.MustCompile(`function\s+([\w_]+(?:\.[\w_]+)+)\s*\((.*?)\)([\s\S]*?)end`).FindStringSubmatch(match)
			pathParts := strings.Split(parts[1], ".")
			tableName := strings.Join(pathParts[:len(pathParts)-1], ".")
			funcName := pathParts[len(pathParts)-1]
			return fmt.Sprintf("%s['%s'] = function(%s)%send", tableName, funcName, parts[2], parts[3])
		})

	code = regexp.MustCompile(`([\w_]+(?:\.[\w_]+)+)\s*=\s*function\s*\((.*?)\)([\s\S]*?)end`).
		ReplaceAllStringFunc(code, func(match string) string {
			parts := regexp.MustCompile(`([\w_]+(?:\.[\w_]+)+)\s*=\s*function\s*\((.*?)\)([\s\S]*?)end`).FindStringSubmatch(match)
			pathParts := strings.Split(parts[1], ".")
			tableName := strings.Join(pathParts[:len(pathParts)-1], ".")
			funcName := pathParts[len(pathParts)-1]
			return fmt.Sprintf("%s['%s'] = function(%s)%send", tableName, funcName, parts[2], parts[3])
		})

	code = regexp.MustCompile(`function\s+([\w_]+(?:\.[\w_]+)*):([\w_]+)\s*\((.*?)\)([\s\S]*?)end`).
		ReplaceAllStringFunc(code, func(match string) string {
			parts := regexp.MustCompile(`function\s+([\w_]+(?:\.[\w_]+)*):([\w_]+)\s*\((.*?)\)([\s\S]*?)end`).FindStringSubmatch(match)
			tablePath := parts[1]
			methodName := parts[2]
			params := parts[3]
			body := parts[4]

			pathParts := strings.Split(tablePath, ".")
			var tableAccess strings.Builder
			tableAccess.WriteString(pathParts[0])
			for _, part := range pathParts[1:] {
				tableAccess.WriteString(fmt.Sprintf("['%s']", part))
			}

			if strings.TrimSpace(params) == "" {
				return fmt.Sprintf("%s['%s'] = function(self)%send", tableAccess.String(), methodName, body)
			}
			return fmt.Sprintf("%s['%s'] = function(self, %s)%send", tableAccess.String(), methodName, params, body)
		})

	return code
}

func replaceFunctionCalls(code string) string {
	code = regexp.MustCompile(`([\w_]+)((\.[\w_]+|\[\'[\w_]+\'\])+)\s*=\s*function`).
		ReplaceAllStringFunc(code, func(match string) string {
			parts := regexp.MustCompile(`([\w_]+)((\.[\w_]+|\[\'[\w_]+\'\])+)\s*=\s*function`).FindStringSubmatch(match)
			base := parts[1]
			fields := parts[2]

			re := regexp.MustCompile(`\.([\w_]+)|\[\'([\w_]+)\'\]`)
			matches := re.FindAllStringSubmatch(fields, -1)

			result := base
			for _, m := range matches {
				if m[1] != "" {
					result += fmt.Sprintf(`['%s']`, m[1])
				} else {
					result += fmt.Sprintf(`['%s']`, m[2])
				}
			}

			return result + " = function"
		})

	code = regexp.MustCompile(`([\w_]+)((?:\.[\w_]+|\[\'[\w_]+\'\])+)(\([^)]*\))`).
		ReplaceAllStringFunc(code, func(match string) string {
			parts := regexp.MustCompile(`([\w_]+)((?:\.[\w_]+|\[\'[\w_]+\'\])+)(\([^)]*\))`).FindStringSubmatch(match)
			base := parts[1]
			fields := parts[2]
			args := parts[3]

			re := regexp.MustCompile(`\.([\w_]+)|\[\'([\w_]+)\'\]`)
			matches := re.FindAllStringSubmatch(fields, -1)

			result := base
			for _, m := range matches {
				if m[1] != "" {
					result += fmt.Sprintf(`['%s']`, m[1])
				} else {
					result += fmt.Sprintf(`['%s']`, m[2])
				}
			}
			return result + args
		})

	code = regexp.MustCompile(`([\w_]+):([\w_]+)(\([^)]*\))`).
		ReplaceAllStringFunc(code, func(match string) string {
			parts := regexp.MustCompile(`([\w_]+):([\w_]+)(\([^)]*\))`).FindStringSubmatch(match)
			table := parts[1]
			method := parts[2]
			args := parts[3]

			if len(args) > 2 {
				return fmt.Sprintf(`%s['%s'](%s, %s`, table, method, table, args[1:])
			}
			return fmt.Sprintf(`%s['%s'](%s%s`, table, method, table, args[1:])
		})

	return code
}
