package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function.
func Remove(request parameter.RemoveRepositoryDataTransferObject) error {
	name := fmt.Sprintf("%s%s%s",
		convention.ToLowerFirstLetter(request.PackageName),
		convention.ToUpperFirstLetter(request.RepositoryName),
		convention.ToUpperFirstLetter(request.Name),
	)
	_ = file.Remove(fmt.Sprintf("internal/repository/model/dto/%s.go", name))
	return nil
}
