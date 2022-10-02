package endpoint

import _ "embed"

var (
	//go:embed resources/endpoint.go.tmpl
	endpointRootTemplate []byte

	//go:embed resources/object.go.tmpl
	endpointObjectTemplate []byte

	//go:embed resources/function.go.tmpl
	endpointFunctionTemplate []byte
)
