package postgres

import _ "embed"

var (
	//go:embed resources/config.go.tmpl
	configTemplate []byte

	//go:embed resources/database.go.tmpl
	databaseTemplate []byte
)
