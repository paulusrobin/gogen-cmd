package dto

type (
	ProjectConfig struct {
		Path   string
		Name   string
		Module string
	}
	AddEndpointParameter struct {
		ProjectConfig
		PackageName  string
		EndpointName string
	}
	AddUsecaseParameter struct {
		ProjectConfig
		PackageName  string
		FunctionName string
	}
	Parameters struct {
		Path         string
		Name         string
		Module       string
		PackageName  string
		EndpointName string
		FunctionName string
	}
)
