# TODO: rewrite dockerfile with lightweight scratch image
FROM golang:1.18-alpine

RUN apk add --update --no-cache git

WORKDIR /go/src/genius-lyrics-parser
COPY . .
RUN go build

ENTRYPOINT ["/go/src/genius-lyrics-parser/genius-lyrics-parser"]