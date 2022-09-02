package dto

import _ "embed"

var (
	//go:embed resources/payload.go.tmpl
	dtoTemplate []byte
)
