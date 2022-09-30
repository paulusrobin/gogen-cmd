package convention

import "strings"

func PackageName(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), "-", "_")
}

func CommandName(name string) string {
	var toBeReplace = map[string]string{
		" ": "-",
		"_": "-",
	}

	name = strings.ToLower(name)
	for key, value := range toBeReplace {
		name = strings.ReplaceAll(name, key, value)
	}
	return name
}

func EndpointName(name string) string {
	return ToUpperFirstLetter(name)
}

func EndpointNameFromFile(name string) string {
	return strings.Split(EndpointName(name), ".")[0]
}

func FunctionName(name string) string {
	return ToUpperFirstLetter(name)
}

func FunctionNameFromFile(name string) string {
	return strings.Split(FunctionName(name), ".")[0]
}

func FunctionsNameFromFile(names []string) []string {
	var response = make([]string, 0)
	for _, name := range names {
		response = append(response, FunctionNameFromFile(name))
	}
	return response
}

func FileName(name string) string {
	return ToLowerFirstLetter(name)
}
