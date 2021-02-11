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

//Command works as exec.Command but add the required OS flags to hide the window for the command
func Command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	return cmd
}

// InstallJava downloads and install Java from AdoptOpenJDK
func InstallJava(filename string) error {
	println("Installing: " + filename + "...")

	//tar -xf OpenJDK8U-jdk_x64_linux_hotspot_8u*.tar.gz
	os.RemoveAll("java")
	os.Mkdir("java", 0755)

	cmd := Command("tar", "-xf", filename, "-C", "java")
	untar, err := cmd.CombinedOutput()
	if err != nil {
		println(string(untar))
		return err
	}

	//searching for the bin folder
	javaBinGlob, _ := filepath.Abs("java/*/bin")
	bins, err := filepath.Glob(javaBinGlob)
	if err != nil {
		return err
	} else if len(bins) == 0 {
		return errors.New("no java bin folder found")
	}

	//finding the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		println("Couldn't find user home directory")
		return err
	}

	//exporting the bin folder to PATH
	path := "\n# added by go-java-launcher\nexport PATH=$PATH:" + bins[0] + "\n"

	FileAppend(homeDir+"/.profile", path)
	FileAppend(homeDir+"/.bashrc", path)

	return nil
}
