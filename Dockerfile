FROM alpine:3.4

RUN apk -U add ca-certificates

ENV APP_ENV dev

EXPOSE 8080

ADD hello-world /bin/hello-world

CMD ["hello-world"]