FROM golang:latest

WORKDIR $GOPATH/src/github.com/kcratie/dirls
COPY . $GOPATH/src/github.com/kcratie/dirls
RUN go build .

ENTRYPOINT ["./dirls"]
