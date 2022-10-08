package convention

import "strings"

var specialCase = map[string]string{"id": "ID"}

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

func FunctionName(name string) string {
	return ToUpperFirstLetter(name)
}

func InterfaceName(name string) string {
	return "I" + FunctionName(name)
}

func FunctionFromFile(name string) string {
	return strings.Split(FunctionName(name), ".")[0]
}

func FileName(name string) string {
	return ToLowerFirstLetter(name)
}
