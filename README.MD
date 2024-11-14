# Go Java Launcher

## Overview

`Go Java Launcher` is a launcher for JAR files that searches for the appropriate JRE version and launches the JAR with that version. It ensures compatibility by downloading and installing the required JRE from AdoptOpenJDK if it is not already installed.

## Problem Statement

Developers often encounter issues when users do not have the required JRE installed or are using an outdated version. This can lead to compatibility problems and bug reports that are difficult to resolve.

## Solution

`Go Java Launcher` addresses these issues by automatically managing the JRE version. It ensures that the correct JRE is used, even if it needs to be downloaded and installed, all without requiring user interaction.

## Getting Started

### Release

Pre-built binaries for Windows and Linux (x64) are available on the [releases](https://github.com/your-repo/releases) page. Note that the Windows release may have a default icon and no app manifest. You can customize these using tools like Resource Hacker, or you can build the launcher from source.

### Building from Source

To build the launcher, you need [Golang](https://golang.org/dl/) and **Make**. Run the following commands:

For Windows:
```
make windows
```

For Linux:
```
make linux
```

If you prefer not to install **Make**, you can manually execute the commands listed in the [Makefile](Makefile).

On Windows, you can customize the launcher icon by replacing [resources/icon.ico](resources/icon.ico).

## Configuration

Configuration is managed via a `go-java.json` file placed in the same directory as the launcher. This JSON file allows you to specify the desired Java version and the name of the JAR file to be launched. A default configuration file is available here: [go-java.json](go-java.json).
