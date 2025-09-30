FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/server/

FROM scratch
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/cmd/server/.env .env
CMD ["./server"]
