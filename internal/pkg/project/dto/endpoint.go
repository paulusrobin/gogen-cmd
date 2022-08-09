package dto

type (
	AddEndpointParameter struct {
		ProjectConfig
		PackageName  string
		EndpointName string
	}
	GenerateUsecase struct {
		ProjectConfig
		PackageName string
	}
	GeneratePackage struct {
		ProjectConfig
		PackageName      string
		UsecaseFunctions []string
	}
)
