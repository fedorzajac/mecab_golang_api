FROM golang:1.20.3-alpine3.17 AS build-env
ADD . /src
# RUN cd /src && go build -o goapp
RUN cd /src && GOOS=linux GOARCH=amd64 go build -ldflags "-w" -o goapp
# GOOS=linux GOARCH=amd64 go build -o myapp-linux-x86_64

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /src/goapp /app/
COPY . .
ENTRYPOINT GIN_MODE=release ./goapp
