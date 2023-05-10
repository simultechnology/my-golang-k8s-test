.PHONY: build

GOOS ?= linux
GOARCH ?= amd64

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o my-golang-k8s-test