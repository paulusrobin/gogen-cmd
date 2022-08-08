package cmd

import _ "embed"

var (
	//go:embed resources/grpc/grpc.go.tmpl
	grpcTemplate []byte

	//go:embed resources/http/http.go.tmpl
	httpTemplate []byte

	//go:embed resources/subscriber/subscriber.go.tmpl
	subscriberTemplate []byte

	//go:embed resources/main.go.tmpl
	mainTemplate []byte
)
