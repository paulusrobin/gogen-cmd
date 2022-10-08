package parameter

type (
	AddDataTransferObject struct {
		ProjectConfigWithPackage
		Name string
		Type string
	}
	RemoveDataTransferObject struct {
		ProjectConfigWithPackage
		Name string
		Type string
	}
	AddRepositoryDataTransferObject struct {
		ProjectConfig
		RepositoryName string
		Name           string
		Type           string
	}
	RemoveRepositoryDataTransferObject struct {
		ProjectConfig
		RepositoryName string
		Name           string
		Type           string
	}
)
