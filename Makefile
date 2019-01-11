VERSION=0.1.0
REVISION := $(shell git rev-parse --short HEAD;)

CLI_EXE=chart-scanner
CLI_PKG=github.com/jdolitsky/chart-scanner/cmd/chart-scanner

.PHONY: build
build: build-linux build-mac build-windows

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -v \
		--ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION)" \
		-o bin/windows/amd64/$(CLI_EXE) $(CLI_PKG)  # windows

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v \
		--ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION)" \
		-o bin/linux/amd64/$(CLI_EXE) $(CLI_PKG)  # linux

build-mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -v \
		--ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION)" \
		-o bin/darwin/amd64/$(CLI_EXE) $(CLI_PKG) # mac osx

.PHONY: clean
clean:
	git status --ignored --short | grep '^!! ' | sed 's/!! //' | xargs rm -rf

.PHONY: get-version
get-version:
	@echo $(VERSION)
