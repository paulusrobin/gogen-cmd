package convention

import "strings"

func PackageName(name string) string {
	return strings.ToLower(name)
}

func EndpointName(name string) string {
	return strings.ToUpper(string(name[0])) + name[1:]
}

func FileName(name string) string {
	return name
}
