BIN=jpp
BUILD_OUTPUT=bin
GO=go

all: clean build

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build: deps
	${GO} build -o ${BUILD_OUTPUT}/${BIN} ./cmd/jpp

test: deps
	${GO} test -v ./...

lintdeps:
	go get -u golang.org/x/lint/golint

lint: lintdeps build
	go vet
	golint -set_exit_status

cross: build crossdeps
	goxz -os=linux,darwin,windows -arch=386,amd64 -n $(BIN) ./cmd/jpp

crossdeps:
	go get github.com/Songmu/goxz/cmd/goxz

clean:
	rm -rf bin goxz vendor
	go clean

.PHONY: deps build test lintdeps lint cross crossdeps clean
