FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./azure-disk-mitigator /azure-disk-mitigator

ENTRYPOINT ["/azure-disk-mitigator"]
