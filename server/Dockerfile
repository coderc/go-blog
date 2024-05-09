FROM golang:1.21.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/server.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/server .

CMD ["./server"]