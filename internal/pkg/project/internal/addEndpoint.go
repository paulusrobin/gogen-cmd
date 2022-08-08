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

	generatedFolders := []dto.ProjectPath{
		{"internal", false},
		{"internal/pkg", false},
		{fmt.Sprintf("internal/pkg/%s", packageFileName), true},
		{fmt.Sprintf("internal/pkg/%s/endpoint", packageFileName), false},
		{fmt.Sprintf("internal/pkg/%s/payload", packageFileName), false},
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/payload/%sEndpoint.go", packageFileName, endpointFileName): string(pkgPayloadEndpointTemplate),
		fmt.Sprintf("internal/pkg/%s/endpoint/%s.go", packageFileName, endpointFileName):        string(pkgEndpointTemplate),
	}

	for _, folder := range generatedFolders {
		generatedPath := path.Join(parameter.Path, folder.Path)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
		if folder.IsPackage {
			if err := GeneratePackage(dto.GeneratePackage{
				ProjectConfig: parameter.ProjectConfig,
				PackageName:   packageFileName,
			}); err != nil {
				return err
			}
		}
	}

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(parameter.Path, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}
