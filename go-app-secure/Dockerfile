FROM golang:1.7.5-alpine

ENV SOURCES /go/src/github.com/rashika/go-app-secure/

COPY . ${SOURCES}

WORKDIR ${SOURCES}

RUN CGO_ENABLED=0 go build -o go-app-secure

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ./go-app-secure
