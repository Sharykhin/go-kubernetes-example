FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/go-hello-world

WORKDIR /go/src/github.com/Sharykhin/go-hello-world

RUN go get .

RUN go install github.com/Sharykhin/go-hello-world

ENTRYPOINT /go/bin/go-hello-world

EXPOSE 3002