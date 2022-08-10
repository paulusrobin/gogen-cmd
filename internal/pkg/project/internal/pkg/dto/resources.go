package dto

import _ "embed"

var (
	//go:embed resources/payload-endpoint.go.tmpl
	endpointTemplate []byte

	//go:embed resources/payload-usecase.go.tmpl
	usecaseTemplate []byte
)
