FROM golang:1.9

ENV APP_ENV dev

EXPOSE 8080

ADD hello-world /bin/hello-world

CMD ["hello-world"]