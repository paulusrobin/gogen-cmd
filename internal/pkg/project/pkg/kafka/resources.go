package kafka

import _ "embed"

var (
	//go:embed resources/config.go.tmpl
	configTemplate []byte

	//go:embed resources/message.go.tmpl
	messageTemplate []byte

	//go:embed resources/kafka.go.tmpl
	kafkaTemplate []byte
)
