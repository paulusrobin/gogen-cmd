package repository

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
	"strings"
)

func generateObject(parameters parameter.ProjectConfigWithRepository) error {
	repositoryPath := convention.PackageName(parameters.RepositoryPath)
	fileOutput := path.Join(parameters.Path, repositoryPath, "root.go")
	if file.Exist(fileOutput) {
		return nil
	}

	// TODO: pass parameter to add file parameter
	return file.Generate(fileOutput, string(repositoryTemplate), map[string]interface{}{})
}

func generateRoot(request parameter.ProjectConfigWithRepository) error {
	packagePath := path.Join(request.Path, "internal/repository")
	fileOutput := path.Join(packagePath, "repositories.go")

	fileNames, err := directory.DirNames(path.Join(packagePath))
	if err != nil {
		return err
	}

	// TODO: get repository name
	//parameters := map[string]interface{}{
	//	"Repositories": []map[string]interface{}{
	//		{
	//			"Name": "",
	//		},
	//	},
	//}

	var usecaseFunctions = make([]string, 0)
	for _, fileName := range fileNames {
		if strings.ToLower(fileName) == "root.go" {
			continue
		}
		usecaseFunctions = append(usecaseFunctions, convention.FunctionNameFromFile(fileName))
	}

	_ = file.Remove(fileOutput)
	return file.Generate(fileOutput, string(interfacesTemplate), map[string]interface{}{
		"ProjectName":      request.Name,
		"ProjectModule":    request.Module,
		"PackageName":      request.PackageName,
		"UsecaseFunctions": usecaseFunctions,
	})
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
