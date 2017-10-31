package frango

import (
	"strconv"
	"strings"
)

/* -----------Int----------- */

func IntToString(dataInt int) string {
	dataString := strconv.Itoa(dataInt)
	return dataString
}

func IntToByteArray(dataInt int) []byte {
	dataByteArray := StringToByteArray(IntToString(dataInt))
	return dataByteArray
}

func IntToBool(dataInt int) bool {
	dataBool := (dataInt != 0)
	return dataBool
}

func IntToUint(dataInt int) uint {
	dataUint := uint(dataInt)
	return dataUint
}

/* -----------Uint----------- */

func UintToInt(dataUint uint) int {
	dataInt := int(dataUint)
	return dataInt
}

func UintToString(dataUint uint) string {
	dataString := strconv.Itoa(UintToInt(dataUint))
	return dataString
}

/* -----------String----------- */

func StringToInt(dataString string) int {
	dataInt, _ := strconv.Atoi(dataString)
	return dataInt
}

func StringToByteArray(dataString string) []byte {
	dataByteArray := []byte(dataString)
	return dataByteArray
}

func StringToBool(dataString string) bool {
	dataBool := !(strings.ToLower(strings.TrimSpace(dataString)) == "false" || strings.TrimSpace(dataString) == "0" || strings.TrimSpace(dataString) == "")
	return dataBool
}

func StringToIntSlice(dataString string, separator string) []int {
	dataIntSlice := StringSliceToIntSlice(StringToStringSlice(dataString, separator))
	return dataIntSlice
}

func StringToStringSlice(dataString string, separator string) []string {
	dataStringSlice := strings.Split(dataString, separator)
	return dataStringSlice
}

func StringToFloat64(dataString string) float64 {
	dataFloat64, _ := strconv.ParseFloat(dataString, 64)
	return dataFloat64
}

/* -----------Byte Array----------- */

func ByteArrayToInt(dataByteArray []byte) int {
	dataInt := StringToInt(ByteArrayToString(dataByteArray))
	return dataInt
}

func ByteArrayToString(dataByteArray []byte) string {
	dataString := string(dataByteArray[:])
	return dataString
}

func ByteArrayToBool(dataByteArray []byte) bool {
	dataBool := StringToBool(ByteArrayToString(dataByteArray))
	return dataBool
}

func ByteArrayToIntSlice(dataByteArray []byte, separator string) []int {
	dataIntSlice := StringToIntSlice(ByteArrayToString(dataByteArray), separator)
	return dataIntSlice
}

func ByteArrayToStringSlice(dataByteArray []byte, separator string) []string {
	dataStringSlice := StringToStringSlice(ByteArrayToString(dataByteArray), separator)
	return dataStringSlice
}

/* -----------Bool----------- */

func BoolToInt(dataBool bool) int {
	var dataInt int
	if dataBool == true {
		dataInt = 1
	} else {
		dataInt = 0
	}
	return dataInt
}

func BoolToString(dataBool bool) string {
	var dataString string
	if dataBool == true {
		dataString = "true"
	} else {
		dataString = "false"
	}
	return dataString
}

func BoolToByteArray(dataBool bool) []byte {
	var dataByteArray []byte
	if dataBool == true {
		dataByteArray = StringToByteArray("true")
	} else {
		dataByteArray = StringToByteArray("false")
	}
	return dataByteArray
}

/* -----------Int Slice----------- */

func IntSliceToString(dataIntSlice []int, separator string) string {
	dataString := strings.Join(IntSliceToStringSlice(dataIntSlice), separator)
	return dataString
}

func IntSliceToByteArray(dataIntSlice []int, separator string) []byte {
	dataByteArray := StringToByteArray(IntSliceToString(dataIntSlice, separator))
	return dataByteArray
}

func IntSliceToStringSlice(dataIntSlice []int) []string {
	dataStringSlice := []string{}
	for i := range dataIntSlice {
		dataStringSlice = append(dataStringSlice, IntToString(dataIntSlice[i]))
	}
	return dataStringSlice
}

/* -----------String Slice----------- */

func StringSliceToString(dataStringSlice []string, separator string) string {
	dataString := strings.Join(dataStringSlice, separator)
	return dataString
}

func StringSliceToByteArray(dataStringSlice []string, separator string) []byte {
	dataByteArray := StringToByteArray(StringSliceToString(dataStringSlice, separator))
	return dataByteArray
}

func StringSliceToIntSlice(dataStringSlice []string) []int {
	dataIntSlice := []int{}
	for i := range dataStringSlice {
		dataIntSlice = append(dataIntSlice, StringToInt(dataStringSlice[i]))
	}
	return dataIntSlice
}
