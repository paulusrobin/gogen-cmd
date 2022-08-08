package dto

type (
	ProjectConfig struct {
		Path   string
		Name   string
		Module string
	}
	ProjectPath struct {
		Path      string
		IsPackage bool
	}
)
