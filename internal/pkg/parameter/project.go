package parameter

type (
	ProjectConfig struct {
		Path   string
		Name   string
		Module string
	}
	ProjectConfigWithPackage struct {
		ProjectConfig
		PackageName string
	}
	ProjectConfigWithRepository struct {
		ProjectConfig
		RepositoryPath string
		PackageName    string
		RepositoryName string
	}
)
