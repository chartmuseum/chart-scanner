FROM golang:1.11-alpine as builder
ADD . /go/src/github.com/jdolitsky/chart-scanner
WORKDIR /go/src/github.com/jdolitsky/chart-scanner
RUN go install github.com/jdolitsky/chart-scanner/cmd/chart-scanner

FROM alpine
RUN apk --update add ca-certificates
COPY --from=builder /go/bin/chart-scanner /bin/chart-scanner
RUN mkdir /workspace
WORKDIR /workspace
ENTRYPOINT  ["/bin/chart-scanner"]
