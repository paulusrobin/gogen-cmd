package parameter

type (
	AddModel struct {
		ProjectConfig
		PackageName string
		ModelName   string
		SkipIfError bool
	}

	RemoveModel struct {
		ProjectConfig
		ModelName string
	}
)
