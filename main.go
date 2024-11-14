package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	helper "github.com/celedev97/go-java-launcher/javahelper"
)

func criticalError(err error) {
	println(err.Error())
	os.Exit(1)
}

const CONFIG string = "go-java.json"

func main() {
	javaVersion := 21
	launch := "./app.jar"

	//reading the configuration file
	bytes, err := os.ReadFile(CONFIG)
	if err == nil {
		//setting a json decoder that pass the numbers as strings (better than having an unknow type)
		decoder := json.NewDecoder(strings.NewReader((string)(bytes)))
		decoder.UseNumber()

		//decoding the json
		var settings map[string]interface{}
		err = decoder.Decode(&settings)
		if err == nil {
			//reading javaVersion
			if val, exists := settings["javaVersion"]; exists {
				javaVersion, err = strconv.Atoi(val.(json.Number).String())
				if err != nil {
					javaVersion = 21
				}
			}
			if val, exists := settings["launch"]; exists {
				launch = val.(string)
			}
		}
	}

	//scanning for the right java version
	java, err := helper.GetJava(javaVersion)
	//if there was an error getting Java i install it then try again
	if err != nil {
		println(err.Error())
		filename, err := helper.DownloadJava(javaVersion)
		if err != nil {
			criticalError(err)
		}
		err = helper.InstallJava(filename, javaVersion)
		if err != nil {
			criticalError(err)
		}
		java, err = helper.GetJava(javaVersion)
		if err != nil {
			criticalError(err)
		}
	}

	println("FOUND JAVA: " + java)

	//get arguments
	args := os.Args[1:]

	err = helper.RunJava(java, launch, args)
	if err != nil {
		criticalError(err)
	}
}
