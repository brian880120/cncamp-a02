FROM golang:1.16 as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -o main .

FROM ubuntu:20.04

WORKDIR /root/

COPY --from=builder /go/src/app/main .

EXPOSE 8090

CMD ["./main"]
