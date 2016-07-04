package main

import (
	"os/exec"
	"path"
	"runtime"
	"os"
)

func main() {
	cmd := exec.Command(path.Join(runtime.GOROOT(), "bin", "go"), append([]string{"vet"}, os.Args[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
