package usecase

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
	"strings"
)

func generateObject(parameters parameter.ProjectConfigWithPackage) error {
	packageFileName := convention.FileName(parameters.PackageName)
	fileOutput := path.Join(parameters.Path, "internal/pkg", packageFileName, "usecase/root.go")
	if file.Exist(fileOutput) {
		return nil
	}

	return file.Generate(fileOutput, string(usecaseObjectTemplate), map[string]interface{}{
		"ProjectName":   parameters.Name,
		"ProjectModule": parameters.Module,
		"PackageName":   parameters.PackageName,
	})
}

func generateRoot(parameters parameter.ProjectConfigWithPackage) error {
	packageFileName := convention.FileName(parameters.PackageName)
	packagePath := path.Join(parameters.Path, "internal/pkg", packageFileName)
	fileOutput := path.Join(packagePath, "usecase.go")

	fileNames, err := directory.FileNames(path.Join(packagePath, "usecase"))
	if err != nil {
		return err
	}

	var usecaseFunctions = make([]string, 0)
	for _, fileName := range fileNames {
		if strings.ToLower(fileName) == "root.go" {
			continue
		}
		usecaseFunctions = append(usecaseFunctions, convention.FunctionFromFile(fileName))
	}

	_ = file.Remove(fileOutput)
	return file.Generate(fileOutput, string(usecaseRootTemplate), map[string]interface{}{
		"ProjectName":      parameters.Name,
		"ProjectModule":    parameters.Module,
		"PackageName":      parameters.PackageName,
		"UsecaseFunctions": usecaseFunctions,
	})
}

// Generate function.
func Generate(request parameter.ProjectConfigWithPackage) error {
	packageFileName := convention.FileName(request.PackageName)
	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/usecase", packageFileName),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generateObject(request)),
		functions.MakeFunc(generateRoot(request)),
	})
}
