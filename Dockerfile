FROM golang:latest

WORKDIR /go/src/api.gateway
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["make", "run"]