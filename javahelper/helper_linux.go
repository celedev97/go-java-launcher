// +build linux
package javahelper

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

// WHERE = Where/which command for current OS
const WHERE = "which"

// OS = Operating System name for AdoptOpenJDK API
const OS = "linux"

// ARCHITECTURE = Architecture name for AdoptOpenJDK API
const ARCHITECTURE = "x64"

// INSTALLER = installer url fragment for AdoptOpenJDK API
const INSTALLER = "binary"

// EXTENSION = Extension for the downloaded file
const EXTENSION = ".tar.gz"

var homeDir string
var javaDir string

func init() {
	homeDir, _ = os.UserHomeDir()
	javaDir = homeDir + "/.javas"
}

//Command works as exec.Command but add the required OS flags to hide the window for the command
func Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

//GetJava get the desired java from the installed ones,
//return an error if something goes wrong or if there is no java
func GetJava(version int) (string, error) {
	installedJavas, _ := whereJava()
	validJavas := filterJavas(version, installedJavas)
	if len(validJavas) > 0 {
		return validJavas[0], nil
	}

	//searching for the bin folder
	bins, err := filepath.Glob(javaDir + "/*/bin/java")
	if err != nil {
		return "", err
	} else if len(bins) != 0 {
		return bins[0], nil
	}

	return "", errors.New("Can't find a suitable java")
}

// InstallJava downloads and extract Java from AdoptOpenJDK
func InstallJava(filename string) error {
	println("Installing: " + filename + "...")

	//tar -xf OpenJDK8U-jdk_x64_linux_hotspot_8u*.tar.gz
	os.Mkdir(javaDir, 0755)

	cmd := Command("tar", "-xf", filename, "-C", javaDir)
	untar, err := cmd.CombinedOutput()
	if err != nil {
		println(string(untar))
		return err
	}

	return nil
}
