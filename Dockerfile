FROM golang:1.13.6-alpine AS build-env
RUN apk update && apk add git make
COPY . /go/src/github.com/capitalonline/cdscloud-controller-manager
RUN cd /go/src/github.com/capitalonline/cdscloud-controller-manager && make container-binary

COPY --from=build-env /cdscloud-controller-manager /cdscloud-controller-manager

ENTRYPOINT ["/cloud-controller-manager"]
