package convention

import "strings"

func ToUpperFirstLetter(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

func ToLowerFirstLetter(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}
