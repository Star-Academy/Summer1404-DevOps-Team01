FROM golang:1.24.5-alpine

WORKDIR /app

COPY go-src/go.mod go-src/go.sum ./

RUN go mod download

COPY go-src/ .

EXPOSE 8080

CMD ["go", "run", "main.go"]