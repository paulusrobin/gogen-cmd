package project

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"path"
	"strings"
)

// IsGolangProject function to validate a path is project root of golang project or not.
// return string project module.
func IsGolangProject(basePath string) (string, bool) {
	pathProject := path.Join(basePath, "go.mod")
	if !file.Exist(pathProject) {
		return "", false
	}

	contents, err := file.Read(pathProject)
	if err != nil {
		return "", false
	}

	lineZero := strings.Split(contents[0], " ")
	return lineZero[len(lineZero)-1], true
}

// Name function to get project name from module name.
func Name(module string) string {
	modules := strings.Split(module, "/")
	return modules[len(module)-1]
}

// ValidateProject function to validate a path is project root of golang project or not.
// return string project name.
// return string project module.
// return error if it is not a golang project.
func ValidateProject(basePath string) (string, string, error) {
	moduleName, isGolangProject := IsGolangProject(basePath)
	if !isGolangProject {
		return "", moduleName, fmt.Errorf("it is not in the golang project folder")
	}
	return Name(moduleName), moduleName, nil
}
