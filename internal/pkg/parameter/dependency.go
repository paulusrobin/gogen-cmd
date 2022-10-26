package parameter

type (
	Dependency    string
	AddDependency struct {
		ProjectConfig
		Dependency Dependency
	}
)

const (
	Postgres Dependency = "postgres"
	Redis    Dependency = "redis"
	Kafka    Dependency = "kafka"
)
