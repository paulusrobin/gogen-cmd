package cmd

import _ "embed"

var (
	//go:embed resources/http/http.go.tmpl
	httpTemplate []byte

	//go:embed resources/main.go.tmpl
	mainTemplate []byte

	//go:embed resources/server.go.tmpl
	serverTemplate []byte
)
