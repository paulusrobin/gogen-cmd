package root

import "os/exec"

// Fmt function to format project modules.
func Fmt() error {
	return exec.Command("gofmt", "-s", "-w", ".").Run()
}
