package usecase

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function to remove an endpoint.
func Remove(request parameter.RemoveUsecase) error {
	_ = file.Remove(fmt.Sprintf("internal/pkg/%s/usecase/%s.go",
		convention.FileName(request.PackageName), convention.FileName(request.FunctionName)))
	return Generate(parameter.ProjectConfigWithPackage{
		ProjectConfig: request.ProjectConfig,
		PackageName:   request.PackageName,
	})
}
