package parameter

type (
	AddCommand struct {
		ProjectConfig
		CommandName string
	}

	RemoveCommand struct {
		ProjectConfig
		CommandName string
	}
)
