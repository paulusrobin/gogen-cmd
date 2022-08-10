package convention

import "strings"

func PackageName(name string) string {
	return strings.ToLower(name)
}

func EndpointName(name string) string {
	return toUpperFirstLetter(name)
}

func EndpointNameFromFile(name string) string {
	return strings.Split(EndpointName(name), ".")[0]
}

func FunctionName(name string) string {
	return toUpperFirstLetter(name)
}

func FunctionNameFromFile(name string) string {
	return strings.Split(FunctionName(name), ".")[0]
}

func FileName(name string) string {
	return name
}
