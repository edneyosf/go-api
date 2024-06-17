FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 5000
CMD ["./main"]