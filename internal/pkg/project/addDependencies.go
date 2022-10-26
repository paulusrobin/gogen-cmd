package project

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/pkg/kafka"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/pkg/postgres"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/pkg/redis"
)

func runPkg(request parameter.AddDependency) error {
	switch request.Dependency {
	case parameter.Postgres:
		return postgres.Add(request)
	case parameter.Redis:
		return redis.Add(request)
	case parameter.Kafka:
		return kafka.Add(request)
	default:
		return fmt.Errorf("invalid dependency")
	}
}

// AddDependencies add external dependencies.
func AddDependencies(request parameter.AddDependency) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(runPkg(request)),
	})
}
