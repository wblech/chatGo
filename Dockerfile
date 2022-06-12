FROM golang:1.18-buster as builder

WORKDIR /srv

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY  . ./chatGo
WORKDIR ./chatGo/src
RUN go build -o ../server
RUN ["chmod", "+x", "../server"]

FROM debian:buster-slim

COPY --from=builder /srv /srv
WORKDIR /srv/chatGo


RUN ["chmod", "+x", "server"]
RUN ["chmod", "+x", "./public/index.html"]
#CMD ["./server"]
CMD tail -f /dev/null