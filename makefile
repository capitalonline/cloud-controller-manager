PKG=github.com/capitalonline/cloud-controller-manager
IMAGE?=registry-bj.capitalonline.net/cck/cloud-controller-manager
VERSION=v0.0.1
LDFLAGS?="-s -w"
.EXPORT_ALL_VARIABLES:

.PHONY: build
build:
	mkdir -p bin
	CGO_ENABLED=0 go build -o bin/cds-ccm ./cmd/

.PHONY: container-binary
container-binary:
	CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -mod vendor -o /cds-ccm ./cmd/
