FROM golang:alpine

WORKDIR /go/src/app
ADD . .

RUN go build -i -v -o app

CMD ["./app"]
