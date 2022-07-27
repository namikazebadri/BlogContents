FROM golang:1.18.4-alpine3.15

ENV GIN_HOME=/app

RUN mkdir -p $GIN_HOME

WORKDIR $GIN_HOME

ADD . $GIN_HOME

RUN go build main.go

EXPOSE 8080

CMD ./main