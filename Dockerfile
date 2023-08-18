FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o userapi cmd/userapi/main.go

CMD ./userapi