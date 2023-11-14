package javahelper

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//DownloadJava downloads a Java installer, and return a string containing the installer location,
//if something goes wrong it returns an error
func DownloadJava(javaVersion int) (string, error) {
	version := strconv.Itoa(javaVersion)
	filename := "adoptopenjdk.jre." + version + "." + ARCHITECTURE + EXTENSION

	if !FileExists(filename) {
		//downloading the jre
		url := "https://api.adoptium.net/v3/" + INSTALLER + "/latest/" + version + "/ga/" + OS + "/" + ARCHITECTURE + "/jre/hotspot/normal/eclipse"
		println("Downloading: " + filename + " from " + url + " ...")

		err := downloadFile(url, filename)
		if err != nil {
			//if there's been an error while downloading i'll remove the downloaded file
			os.Remove(filename)
			return "", err
		}
	}

	return filename, nil
}

//WhereJava run where/which java to search for java installations
func whereJava() ([]string, error) {
	// Executing "where java"
	whereJavaOutput, err := Command(WHERE, "java").CombinedOutput()
	if err != nil {
		return nil, err
	}

	//trimming it to remove eventual extra lines
	stringJavaOutput := strings.Trim(string(whereJavaOutput), " \t\n\r")

	// Looping trough results and checking their versions
	javas := strings.Split(stringJavaOutput, "\n")
	for i, java := range javas {
		//trimming the line again because there are occasionally some \r at the end of the line
		javas[i] = strings.Trim(java, " \t\n\r")
	}

	return javas, nil
}

func filterJavas(version int, javas []string) []string {
	output := []string{}

	for _, java := range javas {
		// Getting the version string
		javaVersionOutput, _ := Command(java, "-version").CombinedOutput()

		// Extracting the version from the string
		re := regexp.MustCompile(`"[0-9\.\_]+"`)
		fullVersion := strings.ReplaceAll(re.FindString(string(javaVersionOutput)), `"`, "")

		//getting only the major version of java as an integer
		majorVersion, _ := strconv.Atoi(strings.Split(fullVersion, ".")[0])
		if version == majorVersion {
			output = append(output, java)
		}
	}
	return output
}

//RunJava is just a shortcut for javaw -jar filename
func RunJava(java string, filename string) error {
	absoluteFileName, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	println(java, " -jar ", absoluteFileName)
	return Command(java, "-jar", absoluteFileName).Start()
}
