package bundler

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func ConvertImageToByteString(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("[ERROR][HEX] Error reading resource %s: %s", path, err.Error())
	}
	byteAsLuaString := []string{}
	for index := range bytes {
		byteAsLuaString = append(byteAsLuaString, "\\x"+strings.ToUpper(hex.EncodeToString(bytes[index:index+1])))
	}
	return fmt.Sprintf("-- LuBu HEX Compressed\nvarname = [[%s]]", strings.Join(byteAsLuaString, ""))
}
