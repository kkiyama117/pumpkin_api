FROM golang:alpine

WORKDIR /go/src/app
ADD . .

RUN go mod download
RUN go build -i

CMD ["./pumpkin_api"]
