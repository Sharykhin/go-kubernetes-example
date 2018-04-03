FROM golang:1.9

ENV APP_ENV dev

ADD . /go/src/github.com/golang/example/outyet

WORKDIR /go/src/github.com/golang/example/outyet

RUN go get .

RUN go install github.com/golang/example/outyet

EXPOSE 8080

ENTRYPOINT /go/bin/outyet