FROM golang:1.13.6-alpine AS build-env
RUN apk update && apk add git make
COPY . /go/src/github.com/capitalonline/cdscloud-controller-manager
RUN cd /go/src/github.com/capitalonline/cdscloud-controller-manager && make container-binary

FROM alpine:3.6
RUN apk update --no-cache && apk add ca-certificates
COPY --from=build-env /cds-ccm /cds-ccm

ENTRYPOINT ["/cds-ccm"]
