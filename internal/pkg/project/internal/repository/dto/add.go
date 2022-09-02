package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
	"strings"
)

// Add function.
func Add(request parameter.AddDataTransferObject) error {
	parameters := map[string]interface{}{
		"Name": convention.FunctionName(request.Name),
		"Type": convention.ToUpperFirstLetter(request.Type),
	}

	packageFileName := convention.FileName(request.PackageName)
	payloadFileName := convention.FileName(request.Name)

	generatedFolders := []string{
		"internal",
		"internal/repository",
	}
	packages := strings.Split(packageFileName, "/")
	packageName := ""
	for _, pkg := range packages {
		packageName = path.Join(packageName, pkg)
		generatedFolders = append(generatedFolders, fmt.Sprintf("internal/repository/%s", packageName))
	}
	generatedFolders = append(generatedFolders, fmt.Sprintf("internal/repository/%s/dto", packageFileName))
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/repository/%s/dto/%s%s.go", packageFileName, payloadFileName, convention.ToUpperFirstLetter(request.Type)): string(dtoTemplate),
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}
