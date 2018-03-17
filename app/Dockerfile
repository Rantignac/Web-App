From golang:alpine


LABEL maintainer="Régis Antignac antignac.regis@gmail.com"

WORKDIR /go/src/app

ADD . /go/src/app

RUN apk add --no-cache git \
    && go get github.com/gorilla/mux \
    && go install app \
    && apk del git

ENTRYPOINT [ "/go/bin/app" ]

EXPOSE 8000