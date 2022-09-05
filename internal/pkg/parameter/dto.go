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
		ProjectConfigWithPackage
		RepositoryName string
		Name           string
		Type           string
	}
	RemoveRepositoryDataTransferObject struct {
		ProjectConfigWithPackage
		RepositoryName string
		Name           string
		Type           string
	}
)
