FROM golang:1.18.4-alpine3.16

ENV GIN_HOME=/app

RUN apk add --update curl && apk add --update git && rm -rf /var/cache/apk/*

RUN mkdir -p $GIN_HOME

WORKDIR $GIN_HOME

ADD . $GIN_HOME

EXPOSE 8080

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air