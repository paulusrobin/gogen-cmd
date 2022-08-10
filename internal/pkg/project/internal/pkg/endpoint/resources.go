package endpoint

import _ "embed"

var (
	//go:embed resources/endpoint.go.tmpl
	endpointTemplate []byte
)
