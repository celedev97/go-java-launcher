// +build windows
package javahelper

import (
	"errors"
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

//GetJava get the desired java from the installed ones,
//return an error if something goes wrong or if there is no java
func GetJava(version int) (string, error) {
	installedJavas, _ := whereJava()
	validJavas := filterJavas(version, installedJavas)
	if len(validJavas) > 0 {
		return validJavas[0], nil
	}
	return "", errors.New("suitable Java not found")
}

// InstallJava downloads and install Java from AdoptOpenJDK
func InstallJava(filename string, version int) error {
	println("Installing: " + filename + "...")
	absFileName, _ := filepath.Abs(filename)

	cmd := Command("cmd.exe", "/C", "msiexec.exe", "/i", absFileName, "/passive")
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
