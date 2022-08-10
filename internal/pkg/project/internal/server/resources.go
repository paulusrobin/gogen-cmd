package server

import _ "embed"

var (
	//go:embed resources/server.go.tmpl
	serverTemplate []byte

	//go:embed resources/http.go.tmpl
	serverHttpTemplate []byte

	//go:embed resources/grpc.go.tmpl
	serverGrpcTemplate []byte

	//go:embed resources/subscriber.go.tmpl
	serverSubscriberTemplate []byte

	//go:embed resources/transport/http/http.go.tmpl
	serverTransportHttpTemplate []byte

	//go:embed resources/transport/grpc/grpc.go.tmpl
	serverTransportGrpcTemplate []byte

	//go:embed resources/transport/pubsub/pubsub.go.tmpl
	serverTransportPubsubTemplate []byte
)
