# Builder
FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app/main.go

# Runner
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/mainrun /app
EXPOSE 8080
CMD /app/mainrun
