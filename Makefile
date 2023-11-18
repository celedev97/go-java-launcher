SHELL := /bin/bash

windows: 
	go get github.com/akavel/rsrc
ifeq ($(OS),Windows_NT)
	go build github.com/akavel/rsrc
	rsrc -ico ./resources/icon.ico
	export GOOS=windows
	export GOARCH=amd64
	go build -o ./bin/go-java-launcher_windows.exe -ldflags -H=windowsgui
endif

linux:
ifeq ($(OS),Windows_NT)
	export GOOS=linux
	export GOARCH=amd64
	go build -o ./bin/go-java-launcher_linux
endif

run:
	go run main.go