package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function.
func Remove(request parameter.RemoveRepositoryDataTransferObject) error {
	_ = file.Remove(fmt.Sprintf("internal/repository/%s/dto/%s.go",
		convention.PackageName(request.RepositoryName), convention.FileName(request.Name)))
	return nil
}
