// +build windows
package main

import (
	"os/exec"
	"syscall"
)

// WHERE = Where/which command for current OS
const WHERE = "where"

// OS = Operating System name for AdoptOpenJDK API
const OS = "windows"

func installJava(architecture string, javaVersion int) (string, error) {
	filename, err := downloadJava(architecture, javaVersion)

	//installing it
	println("Installing: " + exeDir + "\\" + filename + "...")
	cmd := command("cmd.exe", "/C", "msiexec.exe", "/i", exeDir+"\\"+filename, "/passive")
	msiCom, err := cmd.CombinedOutput()
	if err != nil {
		return string(msiCom), err
	}

	println(string(msiCom))
	return getJava(javaVersion)
}

func command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

func progressBarWindow() {

}
