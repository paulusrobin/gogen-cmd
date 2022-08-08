package helper

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
)
