package main

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	var fullPath bytes.Buffer
	// gets the user home path
	fullPath.WriteString(os.Getenv("HOME"))
	// my golang code path
	fullPath.WriteString("/gocode")

	args := []string{"ls", "-a", "-l", "-h", fullPath.String()}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
