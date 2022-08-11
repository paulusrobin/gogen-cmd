package parameter

type (
	AddEndpoint struct {
		ProjectConfig
		PackageName  string
		EndpointName string
	}

	RemoveEndpoint struct {
		ProjectConfig
		PackageName  string
		EndpointName string
	}
)
