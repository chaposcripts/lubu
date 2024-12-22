package bundler

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func ConvertImageToByteString(imagePath string) string {
	bytes, err := os.ReadFile(imagePath)
	if err != nil {
		log.Panicf("[ERROR][IMAGE] Error reading resource %s: %s", imagePath, err.Error())
	}
	byteAsLuaString := []string{}
	for index := range bytes {
		byteAsLuaString = append(byteAsLuaString, "\\x"+strings.ToUpper(hex.EncodeToString(bytes[index:index+1])))
	}
	return fmt.Sprintf("-- LuBu Base85 Compressed\nvarname = [[%s]]", strings.Join(byteAsLuaString, ""))
}
