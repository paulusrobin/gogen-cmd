package repository

import _ "embed"

var (
	//go:embed resources/interfaces.go.tmpl
	interfacesTemplate []byte

	//go:embed resources/repository.go.tmpl
	repositoryTemplate []byte

	//go:embed resources/function.go.tmpl
	repositoryFunctionTemplate []byte
)
