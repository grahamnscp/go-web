FROM golang

ADD ./src /go/src/go-web
RUN go install go-web
CMD /go/bin/go-web

EXPOSE 8080
