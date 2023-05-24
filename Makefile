GO_MODULE := github.com/keybrl/hksr-compass
VERSION ?= 0.0.0-dev
LDFALGS ?= -ldflags="-X 'main.version=$(VERSION)'"

.PHONY: build
build:
	go build -o bin/hksr-compass $(GO_MODULE)

.PHONY: build-release
build-release: clean
	mkdir -p bin/release
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFALGS) -o bin/release/hksr-compass-linux-amd64 $(GO_MODULE)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFALGS) -o bin/release/hksr-compass-linux-arm64 $(GO_MODULE)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFALGS) -o bin/release/hksr-compass-darwin-amd64 $(GO_MODULE)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFALGS) -o bin/release/hksr-compass-darwin-arm64 $(GO_MODULE)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFALGS) -o bin/release/hksr-compass-windows-amd64.exe $(GO_MODULE)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build $(LDFALGS) -o bin/release/hksr-compass-windows-arm64.exe $(GO_MODULE)

.PHONY: clean
clean:
	rm -rf bin
