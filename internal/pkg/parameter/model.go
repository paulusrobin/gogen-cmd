package parameter

type (
	AddModel struct {
		ProjectConfig
		RepositoryName string
		ModelName      string
		SkipIfError    bool
	}

	RemoveModel struct {
		ProjectConfig
		RepositoryName string
		ModelName      string
		SkipIfError    bool
	}
)
