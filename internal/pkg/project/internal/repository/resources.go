package repository

import _ "embed"

var (
	//go:embed resources/repositories.go.tmpl
	repositoryTemplate []byte
)
