// +build windows
package javahelper

import (
	"os/exec"
	"path/filepath"
	"syscall"
)

// WHERE = Where/which command for current OS
const WHERE = "where"

// OS = Operating System name for AdoptOpenJDK API
const OS = "windows"

// ARCHITECTURE = Architecture name for AdoptOpenJDK API
const ARCHITECTURE = "x64"

// INSTALLER = installer url fragment for AdoptOpenJDK API
const INSTALLER = "installer"

// EXTENSION = Extension for the downloaded file
const EXTENSION = ".msi"

// InstallJava downloads and install Java from AdoptOpenJDK
func InstallJava(filename string) error {
	println("Installing: " + filename + "...")

	cmd := Command("cmd.exe", "/C", "msiexec.exe", "/i", filepath.Abs(filename), "/passive")
	msiCom, err := cmd.CombinedOutput()
	if err != nil {
		println(string(msiCom))
		return err
	}

	return nil
}

//Command works as exec.Command but add the required OS flags to hide the window for the command
func Command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

func ProgressBarWindow() {

}
