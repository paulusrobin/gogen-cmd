package model

import _ "embed"

var (
	//go:embed resources/model.go.tmpl
	modelTemplate []byte
)
