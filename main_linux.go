// +build linux
package main

import "os/exec"

// WHERE = Where/which command for current OS
const WHERE = "which"

// OS = Operating System name for AdoptOpenJDK API
const OS = "linux"

func command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	return cmd
}
