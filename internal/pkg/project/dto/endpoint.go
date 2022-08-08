package dto

type (
	AddEndpointParameter struct {
		ProjectConfig
		PackageName  string
		EndpointName string
	}
	GeneratePackage struct {
		ProjectConfig
		PackageName string
	}
)
