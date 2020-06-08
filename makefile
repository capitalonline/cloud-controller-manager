PKG=github.com/capitalonline/cloud-controller-manager
IMAGE?=registry-bj.capitalonline.net/cck/cloud-controller-manager
VERSION=v0.0.1
GIT_COMMIT?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS?="-X ${PKG}/pkg/common.version=${VERSION} -X ${PKG}/pkg/common.gitCommit=${GIT_COMMIT} -X ${PKG}/pkg/common.buildDate=${BUILD_DATE} -s -w"
.EXPORT_ALL_VARIABLES:

.PHONY: build
build:
	mkdir -p bin
	CGO_ENABLED=0 go build -ldflags ${LDFLAGS} -o bin/cdscloud-controller-manager ./cmd/

.PHONY: container-binary
container-binary:
	CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -ldflags ${LDFLAGS} -o /cdscloud-controller-manager ./cmd/
