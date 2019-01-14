FROM golang:1.11-alpine as builder
ADD . /go/src/github.com/jdolitsky/chart-scanner
WORKDIR /go/src/github.com/jdolitsky/chart-scanner
RUN apk add --update make git
RUN make build-linux && mv bin/linux/amd64/chart-scanner /chart-scanner

FROM alpine:3.8
RUN apk --update add ca-certificates
COPY --from=builder /chart-scanner /bin/chart-scanner
RUN mkdir /workspace
WORKDIR /workspace
ENTRYPOINT  ["/bin/chart-scanner"]
