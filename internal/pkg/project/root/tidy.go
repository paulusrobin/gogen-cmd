package root

import "os/exec"

// Tidy function to tidy project modules.
func Tidy() error {
	return exec.Command("go", "mod", "tidy").Run()
}
