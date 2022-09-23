package parameter

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
)

type (
	AddRepository struct {
		ProjectConfig
		RepositoryName string
		PackageName    string
		FunctionName   string
		ModelName      string
	}

	RemoveRepository struct {
		ProjectConfig
		RepositoryName string
		PackageName    string
		FunctionName   string
		ModelName      string
	}

	RepositoryTemplate struct {
		Name      string
		Functions []RepositoryFunctionTemplate
	}

	RepositoryFunctionTemplate struct {
		Name    string
		Package string
	}
)

func (r AddRepository) RepositoryPath() string {
	if r.PackageName != "" {
		packagePath := fmt.Sprintf("internal/repository/%s", convention.PackageName(r.PackageName))
		return fmt.Sprintf("%s/%s", packagePath, convention.PackageName(r.RepositoryName))
	}
	return fmt.Sprintf("internal/repository/%s", convention.PackageName(r.RepositoryName))
}

func (r AddRepository) GeneratedFolders() []string {
	if r.PackageName != "" {
		packagePath := fmt.Sprintf("internal/repository/%s", convention.PackageName(r.PackageName))
		return []string{packagePath, fmt.Sprintf("%s/%s", packagePath, convention.PackageName(r.RepositoryName))}
	}
	return []string{fmt.Sprintf("internal/repository/%s", convention.PackageName(r.RepositoryName))}
}

func (r RemoveRepository) RepositoryPath() string {
	if r.PackageName != "" {
		packagePath := fmt.Sprintf("internal/repository/%s", convention.PackageName(r.PackageName))
		return fmt.Sprintf("%s/%s", packagePath, convention.PackageName(r.RepositoryName))
	}
	return fmt.Sprintf("internal/repository/%s", convention.PackageName(r.RepositoryName))
}
