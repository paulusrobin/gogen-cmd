package config

import _ "embed"

var (
	//go:embed resources/config.go.tmpl
	Template []byte

	//go:embed resources/provider.go.tmpl
	ProviderTemplate []byte

	//go:embed resources/configRuntime.go.tmpl
	RuntimeConfigTemplate []byte
)
