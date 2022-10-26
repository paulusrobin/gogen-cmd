package kafka

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Add add postgres dependency.
func Add(request parameter.AddDependency) error {
	parameters := map[string]interface{}{
		"ProjectName":   request.Name,
		"ProjectModule": request.Module,
	}

	generatedFolders := []string{
		"pkg",
		"pkg/kafka",
	}
	generatedFiles := map[string]string{
		"pkg/kafka/config.go":  string(configTemplate),
		"pkg/kafka/message.go": string(messageTemplate),
		"pkg/kafka/kafka.go":   string(kafkaTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}
