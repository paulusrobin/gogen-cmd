package server

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
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

	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(cfg.Path, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(cfg.Path, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}
