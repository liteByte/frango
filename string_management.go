package frango

import (
	"strings"
)

func FirstLetterToUpper(s string) string {
	bytes := StringToByteArray(s)
	return strings.ToUpper(string(bytes[0])) + string(bytes[1:len(bytes)])
}

func FirstLetterToLower(s string) string {
	bytes := StringToByteArray(s)
	return strings.ToLower(string(bytes[0])) + string(bytes[1:len(bytes)])
}
