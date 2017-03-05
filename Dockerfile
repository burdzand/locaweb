FROM golang:1.7.5

ADD . /go/src/github.com/andersondborges/locaweb
RUN go install github.com/andersondborges/locaweb
ENTRYPOINT /go/bin/locaweb

EXPOSE 9090
