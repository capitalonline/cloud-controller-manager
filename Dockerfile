FROM alpine:3.6

RUN apk add --no-cache ca-certificates

ADD cloud-controller-manager /bin/

CMD ["/bin/cloud-controller-manager"]