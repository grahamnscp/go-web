# go-web - a test repo for raw CI with Docker


## Local test

Build a local image:

using local Dockerfile:

```
FROM golang

ADD ./src /go/src/go-web
RUN go install go-web
CMD /go/bin/go-web

EXPOSE 8080
```

```
  $ docker build -t golang:go-web .
```

Create a container from the image:

```
  $ docker run --name "go-web" -p 8080:8080 -d golang:go-web
```

Quick test:

```
  $ curl http://127.0.0.1:8080/hi
  ho
```

and stop the container:

```
  $ docker stop go-web
```
