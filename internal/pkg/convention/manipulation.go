package convention

import "strings"

func toUpperFirstLetter(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}
