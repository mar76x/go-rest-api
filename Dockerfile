# 1 - Base Go image to build the app binary
FROM golang:1.20 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# 2 - Final lightweight image with only the necessary deps to run the app binary
FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]
