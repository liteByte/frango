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

func RemoveFirstCharacters(s string, n int) string {
	bytes := StringToByteArray(s)
	return string(bytes[n:len(bytes)])
}

func RemoveLastCharacters(s string, n int) string {
	bytes := StringToByteArray(s)
	return string(bytes[0 : len(bytes)-n])
}

func RemoveCharactersUntilSubstring(s, substring string) string {
	substrings := strings.Split(s, substring)
	output := ""

	for i := 1; i < len(substrings); i++ {
		output += substrings[i] + substring
	}

	return RemoveLastCharacters(output, len(substring))
}
