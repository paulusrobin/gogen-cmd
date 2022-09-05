package repository

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"io/fs"
	"path"
)

func generateObject(request parameter.ProjectConfigWithRepository) error {
	repositoryPath := convention.PackageName(request.RepositoryPath)
	fileOutput := path.Join(request.Path, repositoryPath, "root.go")
	if file.Exist(fileOutput) {
		return nil
	}

	return file.Generate(fileOutput, string(repositoryTemplate), map[string]interface{}{
		"PackageName":    convention.PackageName(request.RepositoryName),
		"ProjectModule":  request.Module,
		"RepositoryName": convention.ToUpperFirstLetter(request.RepositoryName),
	})
}

func generateRoot(request parameter.ProjectConfigWithRepository) error {
	packagePath := path.Join(request.Path, "internal/repository")
	fileOutput := path.Join(packagePath, "repositories.go")

	repositories, err := directory.FileNamesWithFilter(packagePath, "^*.*", func(infoPath string, info fs.FileInfo) bool {
		if info.IsDir() && file.Exist(path.Join(infoPath, "root.go")) {
			fmt.Println(infoPath + ": " + info.Name())
			return true
		}
		return false
	})
	if err != nil {
		return err
	}

	parameters := make([]parameter.RepositoryTemplate, 0)
	for _, repository := range repositories {
		parameters = append(parameters, parameter.RepositoryTemplate{
			Name: convention.FunctionName(repository),
		})
	}

	_ = file.Remove(fileOutput)
	return file.Generate(fileOutput, string(interfacesTemplate), map[string]interface{}{"Repositories": parameters})
}

// Generate function.
func Generate(request parameter.ProjectConfigWithRepository) error {
	generatedFolders := []string{
		"internal",
		"internal/repository",
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generateObject(request)),
		functions.MakeFunc(generateRoot(request)),
	})
}
