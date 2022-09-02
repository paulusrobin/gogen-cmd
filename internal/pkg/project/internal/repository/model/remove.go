package model

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function to remove repository model.
func Remove(request parameter.RemoveModel) error {
	_ = file.Remove(fmt.Sprintf("internal/respository/model/%s.go",
		convention.FileName(request.ModelName)))
	return nil
}
