package bundler

import (
	"strings"
)

func RemoveComments(code string) string {
	var result strings.Builder
	inMultilineComment := false
	inString := false
	var stringDelimiter byte
	var commentStart string

	for i := 0; i < len(code); i++ {
		if !inMultilineComment && (code[i] == '"' || code[i] == '\'') {
			if !inString {
				inString = true
				stringDelimiter = code[i]
			} else if code[i] == stringDelimiter {
				inString = false
			}
			result.WriteByte(code[i])
			continue
		}

		if !inString && !inMultilineComment && i+1 < len(code) && code[i] == '-' && code[i+1] == '-' {
			if i+2 < len(code) && code[i+2] == '[' {
				eqCount := 0
				j := i + 3
				for j < len(code) && code[j] == '=' {
					eqCount++
					j++
				}
				if j < len(code) && code[j] == '[' {
					inMultilineComment = true
					commentStart = "--[" + strings.Repeat("=", eqCount) + "["
					i = j
					continue
				}
			}
			for i < len(code) && code[i] != '\n' {
				i++
			}
			if i < len(code) {
				result.WriteByte('\n')
			}
			continue
		}

		if inMultilineComment {
			eqCount := len(commentStart) - 4
			if i+1+eqCount < len(code) {
				endMarker := "]" + strings.Repeat("=", eqCount) + "]"
				if strings.HasPrefix(code[i:], endMarker) {
					inMultilineComment = false
					i += len(endMarker) - 1
					continue
				}
			}
			continue
		}

		if !inMultilineComment {
			result.WriteByte(code[i])
		}
	}

	return result.String()
}
