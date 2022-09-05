package model

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function to remove repository model.
func Remove(request parameter.RemoveModel) error {
	if request.ModelName == "" {
		if request.SkipIfError {
			return nil
		}
		return fmt.Errorf("model name is required")
	}
	_ = file.Remove(fmt.Sprintf("internal/respository/model/%s%s.go",
		convention.FileName(request.RepositoryName), convention.FunctionName(request.ModelName)))
	return nil
}
