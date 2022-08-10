package encoding

import _ "embed"

var (
	//go:embed resources/payload-encoding.go.tmpl
	encodingTemplate []byte
)
