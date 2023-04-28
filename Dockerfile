FROM golang:1.20.3-alpine3.17 AS build-env
ADD . /src
RUN cd /src && go build -o goapp

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /src/goapp /app/
COPY . .
ENTRYPOINT GIN_MODE=release ./goapp
