.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/back-linux

.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/back-darwin


.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/back-windows.exe

.PHONY: build-all
build-all:
	make clean
	mkdir -p ./build
	make build-linux
	make build-darwin
	make build-windows

.PHONY: clean
clean:
	rm -rf ./build