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
)
