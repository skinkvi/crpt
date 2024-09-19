FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY . /app

RUN go build -o crpt .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/crpt .

ENTRYPOINT ["go"]
CMD ["run", "cmd/crpt/main.go"]
