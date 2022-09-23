package server

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Init function to initialize config folder.
func Init(cfg parameter.ProjectConfig) error {
	parameters := map[string]interface{}{
		"ProjectName":   cfg.Name,
		"ProjectModule": cfg.Module,
		"PackageName":   "greeting",
	}
	generatedFolders := []string{
		"internal",
		"internal/server",
		"internal/server/transport",
		"internal/server/transport/grpc",
		"internal/server/transport/http",
		"internal/server/transport/subscriber",
	}
	generatedFiles := map[string]string{
		"internal/server/transport/http/http.go":         string(serverTransportHttpTemplate),
		"internal/server/transport/grpc/grpc.go":         string(serverTransportGrpcTemplate),
		"internal/server/transport/subscriber/pubsub.go": string(serverTransportPubsubTemplate),
		"internal/server/server.go":                      string(serverTemplate),
		"internal/server/grpc.go":                        string(serverGrpcTemplate),
		"internal/server/http.go":                        string(serverHttpTemplate),
		"internal/server/subscriber.go":                  string(serverSubscriberTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(cfg.Path, generatedFolders)),
		functions.MakeFunc(generator.File(cfg.Path, generatedFiles, parameters)),
	})
}
