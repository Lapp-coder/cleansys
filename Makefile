.PHONY: setup-arm64 setup-amd64 build-arm64 build-amd64
.SILENT:

setup-arm64:
	sudo mv ./build/bin/cleansys-arm64 /usr/local/bin/cleansys

setup-amd64:
	sudo mv ./build/bin/cleansys-amd64 /usr/local/bin/cleansys

build-arm64:
	GOOS=darwin GOARCH=arm64 go build -o ./build/bin/cleansys-arm64 main.go

build-amd64:
	GOOS=darwin GOARCH=amd64 go build -o ./build/bin/cleansys-amd64 main.go
