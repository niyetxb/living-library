FROM golang:1.23-alpine

LABEL authors="Akniyet"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
