package internal

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/dto"
	"path"
)

// AddEndpoint function.
func AddEndpoint(parameter dto.AddEndpointParameter) error {
	parameters := map[string]interface{}{
		"ProjectName":   parameter.Name,
		"ProjectModule": parameter.Module,
		"PackageName":   convention.PackageName(parameter.PackageName),
		"EndpointName":  convention.EndpointName(parameter.EndpointName),
	}

	packageFileName := convention.FileName(parameter.PackageName)
	endpointFileName := convention.FileName(parameter.EndpointName)

	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/endpoint", packageFileName),
		fmt.Sprintf("internal/pkg/%s/payload", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/payload/%s.go", packageFileName, endpointFileName):  string(pkgPayloadEndpointTemplate),
		fmt.Sprintf("internal/pkg/%s/endpoint/%s.go", packageFileName, endpointFileName): string(pkgEndpointTemplate),
	}

	for _, folderPath := range generatedFolders {
		if directory.Exist(folderPath) {
			continue
		}
		if err := directory.Make(path.Join(parameter.Path, folderPath)); err != nil {
			return err
		}
	}

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(parameter.Path, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}
