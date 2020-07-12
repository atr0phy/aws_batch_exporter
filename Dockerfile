FROM golang:alpine AS build-env
COPY ./ /opt/workdir/
WORKDIR /opt/workdir/
RUN GOOS=linux GOARCH=amd64 go build -o aws_batch_exporter cmd/aws_batch_exporter.go

FROM alpine:latest
RUN addgroup -S exporterg && adduser -S exporter -G exporterg
USER exporter
COPY --from=build-env /opt/workdir/aws_batch_exporter /opt/aws_batch_exporter
EXPOSE 8080

ENTRYPOINT ["/opt/aws_batch_exporter"]
