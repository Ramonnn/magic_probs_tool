# Start from the official Go image
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o web cmd/web/main.go

EXPOSE 8080
CMD ["./web"]
