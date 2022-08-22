package parameter

type (
	AddRepository struct {
		ProjectConfig
		RepositoryName string
		PackageName    string
		FunctionName   string
	}

	RemoveRepository struct {
		ProjectConfig
		RepositoryName string
		PackageName    string
		FunctionName   string
	}
)
