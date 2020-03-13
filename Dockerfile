FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./azure-disk-mitigator-app /azure-disk-mitigator-app

ENTRYPOINT ["/azure-disk-mitigator-app"]
