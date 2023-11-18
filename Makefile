SHELL := /bin/bash

windows: 
	go get github.com/akavel/rsrc
	go build github.com/akavel/rsrc
	rsrc -ico ./resources/icon.ico
	export GOOS=windows
	export GOARCH=amd64
	go build -o ./bin/go-java-launcher_windows.exe -ldflags -H=windowsgui

linux:
	apt-get install xorg-dev
	export GOOS=linux
	export GOARCH=amd64
	go build -o ./bin/go-java-launcher_linux

run:
	go run main.go