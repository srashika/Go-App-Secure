FROM golang:1.7.5-alpine

ENV SOURCES /go/src/clientsecure/

COPY . ${SOURCES}

WORKDIR ${SOURCES}

RUN CGO_ENABLED=0 go build -o clientsecure

ENTRYPOINT ./clientsecure