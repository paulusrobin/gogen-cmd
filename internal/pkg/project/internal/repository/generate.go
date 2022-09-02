package repository

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Generate function.
func Generate(request parameter.ProjectConfig) error {
	// TODO: get repository name
	parameters := map[string]interface{}{
		"Repositories": []map[string]interface{}{
			{
				"Name": "",
			},
		},
	}
	generatedFolders := []string{
		"internal",
		"internal/repository",
	}
	generatedFiles := map[string]string{
		"internal/repository/repositories.go": string(repositoryTemplate),
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}
