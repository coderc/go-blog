FROM golang:1.21.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -ldflags '-extldflags "-static"' -o server ./cmd/server/server.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config/prod/config.yaml .

ENV GIN_MODE=release
ENV CONFIG=.

EXPOSE 8082

CMD ["./server"]