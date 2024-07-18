FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]

#docker build -t shoe-store-api .
#docker run -p 8080:8080 shoe-store-api
