FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY . .

RUN go build .

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata gcompat

ENV TZ=Europe/Paris

RUN update-ca-certificates

WORKDIR /app

COPY ./templates ./templates

COPY --from=builder /build/main ./main

CMD ["./main"]