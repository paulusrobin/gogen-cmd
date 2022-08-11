package encoding

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function.
func Remove(request parameter.RemoveEndpoint) error {
	_ = file.Remove(fmt.Sprintf("internal/pkg/%s/endpoint/%s.go",
		convention.FileName(request.PackageName), convention.FileName(request.EndpointName)))
	return nil
}
