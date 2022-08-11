package parameter

type (
	AddUsecase struct {
		ProjectConfig
		PackageName  string
		FunctionName string
	}

	RemoveUsecase struct {
		ProjectConfig
		PackageName  string
		FunctionName string
	}
)
