SRC_PKGS=$(shell go list ./... | grep -v vendor)

ifeq ($(strip $(NO_REV)),)
	REV=$(shell git rev-parse --short HEAD)
endif

ifeq ($(BUILD_VERSION),)
	VERSION="1.0.0"
	BUILD_VERSION_NO_REV=$(VERSION)
	ifeq ($(strip $(REV)),)
		BUILD_VERSION=$(VERSION)
	else
		BUILD_VERSION=$(VERSION)-$(REV)
	endif
endif

.PHONY: clean test

all: compile

clean:
	go clean ./src/...
	rm -rf dist

compile:
	CGO_ENABLED=0 go build -ldflags "-X main.appVersion=$(BUILD_VERSION)" .

dist:
	CGO_ENABLED=0 go build -ldflags "-X main.appVersion=$(BUILD_VERSION)" -o dist/golang-mud

test:
	set -e;
	for pkg in $(SRC_PKGS); \
	do \
		go test -v $$pkg; \
	done
