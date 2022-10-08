package parameter

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
)

type (
	AddRepository struct {
		ProjectConfig
		RepositoryName string
		FunctionName   string
		ModelName      string
	}

	RemoveRepository struct {
		ProjectConfig
		RepositoryName string
		FunctionName   string
		ModelName      string
	}
)

func (r AddRepository) RepositoryPath() string {
	return fmt.Sprintf("internal/repository/%s", convention.PackageName(r.RepositoryName))
}

func (r AddRepository) FunctionPath() string {
	return fmt.Sprintf("%s/%s.go", r.RepositoryPath(), convention.FileName(r.FunctionName))
}

func (r AddRepository) GeneratedFolders() []string {
	return []string{r.RepositoryPath()}
}

func (r RemoveRepository) RepositoryPath() string {
	return fmt.Sprintf("internal/repository/%s", convention.PackageName(r.RepositoryName))
}

func (r RemoveRepository) FunctionPath() string {
	return fmt.Sprintf("%s/%s.go", r.RepositoryPath(), convention.FileName(r.FunctionName))
}
