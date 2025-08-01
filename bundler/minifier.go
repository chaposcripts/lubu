package bundler

import (
	"regexp"
	"strings"
)

func MinifyCode(code string) string {
	blockCommentRegex := regexp.MustCompile(`--\[\[.*?\]\]`)
	code = blockCommentRegex.ReplaceAllString(code, "")
	lines := strings.Split(code, "\n")
	var result strings.Builder
	for _, line := range lines {
		if idx := strings.Index(line, "--"); idx >= 0 {
			line = line[:idx]
		}
		line = strings.TrimSpace(line)
		if line != "" {
			result.WriteString(line + " ")
		}
	}
	uglified := result.String()
	uglified = regexp.MustCompile(`\s*([=+\-*/%^#<>~;:{},])\s*`).ReplaceAllString(uglified, "$1")
	uglified = regexp.MustCompile(`\s+`).ReplaceAllString(uglified, " ")
	uglified = regexp.MustCompile(`([({[])\s+`).ReplaceAllString(uglified, "$1")
	uglified = regexp.MustCompile(`\s+([)}])`).ReplaceAllString(uglified, "$1")
	return uglified
}
