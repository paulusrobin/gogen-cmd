package cmd

import _ "embed"

var (
	//go:embed resources/grpc/grpc.go.tmpl
	grpcTemplate string

	//go:embed resources/http/http.go.tmpl
	httpTemplate string

	//go:embed resources/subscriber/subscriber.go.tmpl
	subscriberTemplate string

	//go:embed resources/main.go.tmpl
	mainTemplate string
)
