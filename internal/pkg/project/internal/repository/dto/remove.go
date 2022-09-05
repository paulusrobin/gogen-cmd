package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function.
func Remove(request parameter.RemoveDataTransferObject) error {
	_ = file.Remove(fmt.Sprintf("internal/repository/%s/dto/%s%s.go",
		convention.FileName(request.PackageName),
		convention.FileName(request.Name),
		convention.ToUpperFirstLetter(request.Type)))
	return nil
}
