package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/dto"
	"path"
	"strings"
)

// Generate function to generate internal folder files.
func Generate(cfg dto.ProjectConfig) error {
	parameters := map[string]interface{}{
		"ProjectName":   cfg.Name,
		"ProjectModule": cfg.Module,
		"PackageName":   "greeting",
	}
	generatedFolders := []string{
		"internal",
		"internal/config",
		"internal/pkg",
		"internal/pkg/greeting",
		"internal/pkg/greeting/endpoint",
		"internal/pkg/greeting/payload",
		"internal/pkg/greeting/usecase",
		"internal/repository",
		"internal/repository/model",
		"internal/server",
		"internal/server/transport",
		"internal/server/transport/grpc",
		"internal/server/transport/http",
		"internal/server/transport/subscriber",
	}
	generatedFiles := map[string]string{
		"internal/config/config.go":                      string(configTemplate),
		"internal/config/provider.go":                    string(configProviderTemplate),
		"internal/pkg/greeting/endpoint/greeting.go":     string(greetingEndpointTemplate),
		"internal/pkg/greeting/payload/greeting.go":      string(greetingPayloadTemplate),
		"internal/pkg/greeting/usecase/greeting.go":      string(greetingUsecaseFunctionTemplate),
		"internal/pkg/greeting/usecase/root.go":          string(pkgUsecaseRootTemplate),
		"internal/pkg/greeting/usecase.go":               string(greetingUsecaseTemplate),
		"internal/server/transport/http/http.go":         string(serverTransportHttpTemplate),
		"internal/server/transport/grpc/grpc.go":         string(serverTransportGrpcTemplate),
		"internal/server/transport/subscriber/pubsub.go": string(serverTransportPubsubTemplate),
		"internal/server/server.go":                      string(serverTemplate),
		"internal/server/grpc.go":                        string(serverGrpcTemplate),
		"internal/server/http.go":                        string(serverHttpTemplate),
		"internal/server/subscriber.go":                  string(serverSubscriberTemplate),
	}

	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(cfg.Path, cfg.Name, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(cfg.Path, cfg.Name, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}

// GeneratePackage function to generate new internal pkg folder files.
func GeneratePackage(parameters dto.GeneratePackage) error {
	packagePath := path.Join(parameters.Path, "internal/pkg", parameters.PackageName)
	fileOutput := path.Join(packagePath, "usecase.go")

	fileNames, err := directory.FileNames(path.Join(packagePath, "usecase"))
	if err != nil {
		return err
	}

	var usecaseFunctions = make([]string, 0)
	for _, fileName := range fileNames {
		if strings.ToLower(fileName) == "root.go" {
			continue
		}
		usecaseFunctions = append(usecaseFunctions, convention.FunctionName(fileName))
	}

	_ = file.Remove(fileOutput)
	if err := file.Generate(fileOutput, string(pkgUsecaseTemplate),
		map[string]interface{}{
			"ProjectName":      parameters.Name,
			"ProjectModule":    parameters.Module,
			"PackageName":      parameters.PackageName,
			"UsecaseFunctions": usecaseFunctions,
		}); err != nil {
		return err
	}
	return nil
}

// GenerateUsecase function to generate new internal pkg usecase folder files.
func GenerateUsecase(parameters dto.GenerateUsecase) error {
	fileOutput := path.Join(parameters.Path, "internal/pkg", parameters.PackageName, "usecase/root.go")
	if file.Exist(fileOutput) {
		return nil
	}

	return file.Generate(fileOutput, string(pkgUsecaseRootTemplate), map[string]interface{}{
		"ProjectName":   parameters.Name,
		"ProjectModule": parameters.Module,
		"PackageName":   parameters.PackageName,
	})
}
