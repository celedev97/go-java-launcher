build:
	go get github.com/akavel/rsrc
	rsrc -ico ./resources/icon.ico
	go build -o ./bin/go-java-launcher.exe -ldflags -H=windowsgui

run:
	go run main.go