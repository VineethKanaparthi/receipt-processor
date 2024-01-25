# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -v -o receipt-processor-webservice ./cmd

EXPOSE 8080

CMD ["./receipt-processor-webservice"]
