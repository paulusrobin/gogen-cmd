package parameter

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
)
