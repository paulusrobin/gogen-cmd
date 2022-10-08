package repository

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"io/fs"
	"path"
	"strings"
)

func generateRoot(request parameter.ProjectConfigWithRepository) error {
	packagePath := path.Join(request.Path, "internal/repository", request.RepositoryName)
	fileOutput := path.Join(packagePath, "repository.go")

	repositoryInterfaces, err := directory.FileNamesWithFilter(
		packagePath, directory.AllFilter, func(infoPath string, info fs.FileInfo) bool {
			if info.IsDir() || strings.Contains(infoPath, "dto") || info.Name() == "repository.go" {
				return false
			}
			return true
		})
	if err != nil {
		return err
	}

	var repositoryFunctions = make([]string, 0)
	for _, repositoryInterface := range repositoryInterfaces {
		repositoryFunctions = append(repositoryFunctions, "I"+convention.FunctionFromFile(repositoryInterface))
	}

	_ = file.Remove(fileOutput)
	return file.Generate(fileOutput, string(repositoryTemplate), map[string]interface{}{
		"PackageName": convention.PackageName(request.RepositoryName),
		"Functions":   repositoryFunctions,
	})
}

// Generate function.
func Generate(request parameter.ProjectConfigWithRepository) error {
	generatedFolders := []string{
		"internal",
		"internal/repository",
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generateRoot(request)),
	})
}
