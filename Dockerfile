FROM golang:1.22.5-alpine

WORKDIR /app

COPY . .

RUN go mod download

ENTRYPOINT ["go", "run", "./cmd/crpt/main.go"]

