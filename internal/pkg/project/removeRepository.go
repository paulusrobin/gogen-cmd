package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository"
)

// RemoveRepository function to remove a repository.
func RemoveRepository(request parameter.RemoveRepository) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(repository.Remove(request)),
	})
}
