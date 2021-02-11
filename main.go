package main

import "github.com/fcdev/go-java-launcher/platform"

func main() {
	javaVersion := 11
	java, err := platform.GetJava(javaVersion)
	if err != nil {
		println(err.Error())
		java, err = platform.InstallJava("x64", javaVersion)
		if err != nil {
			println(java)
			println(err.Error())
			platform.Pause()
			return
		}
	}

	println(java)

	err = platform.RunJava(java, "./app.jar")
	if err != nil {
		println(err.Error())
		platform.Pause()
	}
}
