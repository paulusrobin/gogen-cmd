package cmd

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"io/fs"
	"path"
)

// Generate function.
func Generate(request parameter.ProjectConfig) error {
	generatedFolders := []string{"cmd"}
	packagePath := path.Join(request.Path, "cmd")
	commands, err := directory.FileNamesWithFilter(packagePath, directory.AllFilter,
		func(infoPath string, info fs.FileInfo) bool {
			if info.IsDir() && infoPath != packagePath {
				return true
			}
			return false
		})
	if err != nil {
		return err
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(func() error {
			_ = file.Remove(path.Join(request.Path, "cmd", "main.go"))
			return nil
		}()),
		functions.MakeFunc(file.Generate(
			path.Join(request.Path, "cmd", "main.go"),
			string(mainTemplate),
			map[string]interface{}{
				"ProjectModule": request.Module,
				"Commands":      commands,
			},
		)),
	})
}
