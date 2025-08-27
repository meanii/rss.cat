# Dockerfile for RSS.cat Telegram Bot
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o rsscat main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/rsscat .
COPY --from=builder /app/rss.cat.db .
ENV TOKEN=""
CMD ["./rsscat"]
