package internal

import _ "embed"

var (
	// ==================================
	//  Config
	// ==================================

	//go:embed resources/config/config.go.tmpl
	configTemplate []byte

	//go:embed resources/config/provider.go.tmpl
	configProviderTemplate []byte

	//go:embed resources/config/configRuntime.go.tmpl
	configRuntimeTemplate []byte

	// ==================================
	//  Server
	// ==================================

	//go:embed resources/server/server.go.tmpl
	serverTemplate []byte

	//go:embed resources/server/http.go.tmpl
	serverHttpTemplate []byte

	//go:embed resources/server/grpc.go.tmpl
	serverGrpcTemplate []byte

	//go:embed resources/server/subscriber.go.tmpl
	serverSubscriberTemplate []byte

	//go:embed resources/server/transport/http/http.go.tmpl
	serverTransportHttpTemplate []byte

	//go:embed resources/server/transport/grpc/grpc.go.tmpl
	serverTransportGrpcTemplate []byte

	//go:embed resources/server/transport/pubsub/pubsub.go.tmpl
	serverTransportPubsubTemplate []byte

	// ==================================
	//  Pkg
	// ==================================

	//go:embed resources/pkg/endpoint/endpoint.go.tmpl
	pkgEndpointTemplate []byte

	//go:embed resources/pkg/payload/payload-endpoint.go.tmpl
	pkgPayloadEndpointTemplate []byte

	//go:embed resources/pkg/payload/payload-usecase.go.tmpl
	pkgPayloadUsecaseTemplate []byte

	//go:embed resources/pkg/usecase.go.tmpl
	pkgUsecaseTemplate []byte

	//go:embed resources/pkg/usecase/root.go.tmpl
	pkgUsecaseRootTemplate []byte

	//go:embed resources/pkg/usecase/function.go.tmpl
	pkgUsecaseFunctionTemplate []byte

	// ==================================
	//  Greeting
	// ==================================

	//go:embed resources/pkg/endpoint/greeting.go.tmpl
	greetingEndpointTemplate []byte

	//go:embed resources/pkg/payload/greeting.go.tmpl
	greetingPayloadTemplate []byte

	//go:embed resources/pkg/greeting.go.tmpl
	greetingUsecaseTemplate []byte

	//go:embed resources/pkg/usecase/greeting.go.tmpl
	greetingUsecaseFunctionTemplate []byte
)
