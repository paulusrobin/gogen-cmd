package encoding

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
)

// Add function.
func Add(request parameter.AddEndpoint) error {
	parameters := map[string]interface{}{
		"ProjectName":   request.Name,
		"ProjectModule": request.Module,
		"PackageName":   convention.PackageName(request.PackageName),
		"EndpointName":  convention.EndpointName(request.EndpointName),
	}

	packageFileName := convention.FileName(request.PackageName)
	ecnodingFileName := convention.FileName(request.EndpointName)

	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/encoding", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/encoding/%s.go", packageFileName, ecnodingFileName): string(encodingTemplate),
	}

	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(request.Path, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(request.Path, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}
