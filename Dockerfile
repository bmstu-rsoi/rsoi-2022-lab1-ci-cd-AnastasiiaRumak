FROM golang:1.19 AS build

WORKDIR /ap

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./ap ./cmd/app/main.go

EXPOSE 8080

CMD ./ap