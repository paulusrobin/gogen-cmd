package root

import _ "embed"

var (
	//go:embed resources/.env.sample
	envSampleTemplate []byte

	//go:embed resources/.gitignore
	gitIgnoreTemplate []byte

	//go:embed resources/docker-compose.yaml
	dockerComposeTemplate []byte

	//go:embed resources/Dockerfile
	dockerfileTemplate []byte

	//go:embed resources/Dockerfile-builder
	dockerfileBuilderTemplate []byte

	//go:embed resources/go.mod.tmpl
	goModTemplate []byte

	//go:embed resources/Makefile
	makefileTemplate []byte

	//go:embed resources/README.md
	readmeTemplate []byte
)
