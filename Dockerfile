FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o crpt ./cmd/crpt

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/crpt .

COPY config.yaml .

ENTRYPOINT ["./crpt"]
