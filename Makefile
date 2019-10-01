BIN=jpp
BUILD_OUTPUT=bin

all: clean build

deps:
	go mod vendor

build: deps
	go build -o ${BUILD_OUTPUT}/${BIN} ./cmd/jpp

test: deps
	go test -v ./...

# GO111MODULE=off for dropping devtools from go.mod
# https://github.com/golang/go/issues/30515
lintdeps:
	GO111MODULE=off go get -u golang.org/x/lint/golint

lint: lintdeps build
	go vet
	golint -set_exit_status

cross: build crossdeps
	goxz -os=linux,darwin,windows -arch=386,amd64 -n $(BIN) ./cmd/jpp

crossdeps:
	GO111MODULE=off go get github.com/Songmu/goxz/cmd/goxz

clean:
	rm -rf bin goxz vendor
	go clean

.PHONY: deps build test lintdeps lint cross crossdeps clean
