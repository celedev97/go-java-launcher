package main

import (
	"bufio"
	"os"

	helper "github.com/fcdev/go-java-launcher/javahelper"
)

func criticalError(err error) {
	println(err.Error())

	println("Press 'Enter' to close...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	os.Exit(1)
}

func main() {
	javaVersion := 11
	java, err := helper.GetJava(javaVersion)
	//if there was an error getting Java i close it then run it again
	if err != nil {
		println(err.Error())
		filename, err := helper.DownloadJava(javaVersion)
		if err != nil {
			criticalError(err)
		}
		err = helper.InstallJava(filename)
		if err != nil {
			criticalError(err)
		}
		java, err = helper.GetJava(javaVersion)
		if err != nil {
			criticalError(err)
		}
	}

	println("FOUND JAVA: " + java)

	err = helper.RunJava(java, "./app.jar")
	if err != nil {
		criticalError(err)
	}
}
