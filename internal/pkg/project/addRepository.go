package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository"
)

// AddRepository function to add a repository.
func AddRepository(request parameter.AddRepository) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(repository.Add(request)),
	})
}
