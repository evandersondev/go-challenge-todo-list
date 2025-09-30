FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server/

FROM scratch
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/cmd/server/.env .env
CMD ["./server"]
