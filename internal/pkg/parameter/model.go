package parameter

type (
	AddModel struct {
		ProjectConfig
		ModelName   string
		SkipIfError bool
	}

	RemoveModel struct {
		ProjectConfig
		ModelName   string
		SkipIfError bool
	}
)
