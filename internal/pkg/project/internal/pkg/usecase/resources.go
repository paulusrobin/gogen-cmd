package usecase

import _ "embed"

var (
	//go:embed resources/usecase.go.tmpl
	usecaseRootTemplate []byte

	//go:embed resources/object.go.tmpl
	usecaseObjectTemplate []byte

	//go:embed resources/function.go.tmpl
	usecaseFunctionTemplate []byte
)
