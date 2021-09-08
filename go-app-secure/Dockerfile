FROM golang:1.7.5-alpine

ENV SOURCES /go/src/github.com/rashika/go-app/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT go-app
