package redis

import _ "embed"

var (
	//go:embed resources/config.go.tmpl
	configTemplate []byte

	//go:embed resources/cache.go.tmpl
	cacheTemplate []byte
)
