package main

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//getting current folder path for comodity
var exeDir string = getExeDir()

func getExeDir() string {
	exe, _ := os.Executable()
	exeDir, _ := filepath.Abs(path.Dir(exe))
	return exeDir
}

func pause() {
	println("Press 'Enter' to close...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

//Download the file at url to the filepath,
//return an error if something goes wrong otherwise return nil
func downloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//Download a Java installer
//return a string containing the installer location, if something goes wrong it returns an error
func downloadJava(architecture string, javaVersion int) (string, error) {
	version := strconv.Itoa(javaVersion)
	url := "https://api.adoptopenjdk.net/v3/installer/latest/" + version + "/ga/windows/" + architecture + "/jre/hotspot/normal/adoptopenjdk"
	filename := "adoptopenjdk.jre." + version + "." + architecture + ".msi"

	//downloading the jre
	println("Downloading: " + filename + " from " + url + " ...")
	err := downloadFile(url, filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// return the path to the java executable for the desired version of the jre
func getJava(version int) (string, error) {
	// Executing "where java"
	whereJavaOutput, err := command(WHERE, "java").CombinedOutput()
	if err != nil {
		return "", err
	}

	//trimming it to remove eventual extra lines
	stringJavaOutput := strings.Trim(string(whereJavaOutput), " \t\n\r")

	// Looping trough results and checking their versions
	javas := strings.Split(stringJavaOutput, "\n")
	for _, java := range javas {
		//trimming the line again because there are occasionally some \r at the end of the line
		java = strings.Trim(java, " \t\n\r")

		// Getting the version string
		javaVersionOutput, _ := command(java, "-version").CombinedOutput()

		// Extracting the version from the string
		re := regexp.MustCompile(`"[0-9\.\_]+"`)
		fullVersion := strings.ReplaceAll(re.FindString(string(javaVersionOutput)), `"`, "")

		//getting only the major version of java as an integer
		majorVersion, _ := strconv.Atoi(strings.Split(fullVersion, ".")[0])
		if version == majorVersion {
			return strings.Replace(java, "java.exe", "javaw.exe", -1), nil
		}
	}
	return "", errors.New("Java " + strconv.Itoa(version) + " not found!")
}

func main() {
	javaVersion := 11
	java, err := getJava(javaVersion)
	if err != nil {
		println(err.Error())
		java, err = installJava("x64", javaVersion)
		if err != nil {
			println(java)
			println(err.Error())
			pause()
			return
		}
	}

	println(java)

	err = command(java, "-jar", exeDir+"/app.jar").Start()
	if err != nil {
		println(err.Error())
		pause()
	}
}
