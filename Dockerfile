FROM golang:1.16 as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8090

CMD ["make", "run"]
