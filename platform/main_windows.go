// +build windows
package platform

import (
	"os/exec"
	"syscall"
)

// WHERE = Where/which command for current OS
const WHERE = "where"

// OS = Operating System name for AdoptOpenJDK API
const OS = "windows"

// InstallJava downloads and install Java from AdoptOpenJDK
func InstallJava(architecture string, javaVersion int) (string, error) {
	filename, err := DownloadJava(architecture, javaVersion)

	//installing it
	println("Installing: " + exeDir + "\\" + filename + "...")
	cmd := Command("cmd.exe", "/C", "msiexec.exe", "/i", exeDir+"\\"+filename, "/passive")
	msiCom, err := cmd.CombinedOutput()
	if err != nil {
		return string(msiCom), err
	}

	println(string(msiCom))
	return GetJava(javaVersion)
}

//Command works as exec.Command but add the required OS flags to hide the window for the command
func Command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

func ProgressBarWindow() {

}
