SHELL := /bin/bash

windows: 
	go get github.com/akavel/rsrc
ifeq ($(OS),Windows_NT)
	go build github.com/akavel/rsrc
	rsrc -ico ./resources/icon.ico
	SET GOOS=windows&& SET GOARCH=amd64&& go build -o ./bin/go-java-launcher_windows.exe -ldflags -H=windowsgui
else
	go build github.com/akavel/rsrc
	ls -l
	./rsrc -ico ./resources/icon.ico
	env GOOS=windows GOARCH=amd64 go build -o ./bin/go-java-launcher_windows.exe -ldflags -H=windowsgui
endif

linux:
ifeq ($(OS),Windows_NT)
	SET GOOS=linux&& SET GOARCH=amd64&& go build -o ./bin/go-java-launcher_linux
else
	env GOOS=linux GOARCH=amd64 go build -o ./bin/go-java-launcher_linux
endif

run:
	go run main.go