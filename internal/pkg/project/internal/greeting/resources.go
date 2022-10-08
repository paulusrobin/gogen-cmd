package greeting

import _ "embed"

var (
	//go:embed resources/dto/greeting.go.tmpl
	greetingDataTransferObjectTemplate []byte

	//go:embed resources/encoding/greeting.go.tmpl
	greetingDataTransferObjectEncodingTemplate []byte

	//go:embed resources/endpoint/endpoint.go.tmpl
	greetingEndpointTemplate []byte

	//go:embed resources/endpoint/object.go.tmpl
	greetingEndpointObjectTemplate []byte

	//go:embed resources/endpoint/greeting.go.tmpl
	greetingEndpointFunctionTemplate []byte

	//go:embed resources/usecase/usecase.go.tmpl
	greetingUsecaseTemplate []byte

	//go:embed resources/usecase/object.go.tmpl
	greetingUsecaseObjectTemplate []byte

	//go:embed resources/usecase/greeting.go.tmpl
	greetingUsecaseFunctionTemplate []byte
)
